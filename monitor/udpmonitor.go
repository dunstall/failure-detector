package monitor

import (
	"log"
	"net"
	"strconv"
	"time"
)

// TODO(AD) Start with broadcast - later look at multicast (as broadcast cant
// go via routers.
type UDPMonitor struct {
	id         NodeID
	lis        net.PacketConn
	done       chan bool
	broadcast  *net.UDPAddr
	heartbeats chan Heartbeat
	rate       time.Duration
}

func NewUDPMonitor(id NodeID, broadcastIP string, port uint16, rate time.Duration) (Monitor, error) {
	conn, err := net.ListenPacket("udp", listenAddr(port))
	if err != nil {
		return nil, err
	}
	broadcast, err := net.ResolveUDPAddr("udp4", broadcastAddr(broadcastIP, port))
	if err != nil {
		return nil, err
	}
	m := &UDPMonitor{id: id, lis: conn, done: make(chan bool), broadcast: broadcast, heartbeats: make(chan Heartbeat), rate: rate}
	go m.run()
	return m, nil
}

func (m *UDPMonitor) Heartbeats() <-chan Heartbeat {
	return m.heartbeats
}

func (m *UDPMonitor) Close() error {
	close(m.done)
	return m.lis.Close()
}

func (m *UDPMonitor) run() {
	for {
		select {
		case <-m.done:
			return
		case <-time.After(m.rate):
		}

		id, err := m.recv()
		if err != nil {
			log.Println("recv error:", err)
		}
		m.heartbeats <- NewHeartbeat(id)

		if err := m.ping(); err != nil {
			log.Println("ping error:", err)
		}
	}
}

func (m *UDPMonitor) recv() (NodeID, error) {
	var id NodeID
	b := make([]byte, 4)
	m.lis.SetReadDeadline(time.Now().Add(time.Millisecond * 50))
	n, _, err := m.lis.ReadFrom(b)
	if err != nil {
		return id, err
	}
	err = id.Decode(b[:n])
	return id, err
}

func (m *UDPMonitor) ping() error {
	m.lis.SetWriteDeadline(time.Now().Add(time.Millisecond * 50))
	_, err := m.lis.WriteTo(m.id.Encode(), m.broadcast)
	return err
}

func listenAddr(port uint16) string {
	return net.JoinHostPort("", strconv.Itoa(int(port)))
}

func broadcastAddr(broadcastIP string, port uint16) string {
	return net.JoinHostPort(broadcastIP, strconv.Itoa(int(port)))
}
