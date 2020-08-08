package detector

import (
	"testing"

	"github.com/dunstall/failure-detector/accrual"
	"github.com/dunstall/failure-detector/monitor"
)

type FakeMonitor struct {
	heartbeats chan monitor.Heartbeat
}

func NewFakeMonitor() FakeMonitor {
	return FakeMonitor{heartbeats: make(chan monitor.Heartbeat)}
}

func (m *FakeMonitor) Heartbeats() <-chan monitor.Heartbeat {
	return m.heartbeats
}

func (m *FakeMonitor) Close() error {
	close(m.heartbeats)
	return nil
}

func TestDetectorInit(t *testing.T) {
	m := NewFakeMonitor()
	d := NewDetector(&m, 5)
	defer d.Close()
	if _, err := d.Phi(100, monitor.NodeID(5)); err != ErrNotFound {
		t.Errorf("expected node not found")
	}
}

func TestDetectorReceiveHeartbeat(t *testing.T) {
	m := NewFakeMonitor()
	d := NewDetector(&m, 5)
	defer d.Close()

	id := monitor.NodeID(4)

	m.heartbeats <- monitor.Heartbeat{NodeID: id, Received: 100}
	m.heartbeats <- monitor.Heartbeat{NodeID: id, Received: 150}
	m.heartbeats <- monitor.Heartbeat{NodeID: id, Received: 250}

	expected := accrual.Suspicion{Phi: 2.869699, NSamples: 3}
	var time uint64 = 400
	suspicion, err := d.Phi(time, id)
	if err != nil {
		t.Error(err)
	}
	if !suspicion.Equal(expected) {
		t.Errorf("acc.Phi(%d) != %#v, actual %#v", time, expected, suspicion)
	}

	m.heartbeats <- monitor.Heartbeat{NodeID: id, Received: 450}
	m.heartbeats <- monitor.Heartbeat{NodeID: id, Received: 550}

	expected = accrual.Suspicion{Phi: 9.533335722, NSamples: 5}
	time = 1000
	suspicion, err = d.Phi(time, id)
	if err != nil {
		t.Error(err)
	}
	if !suspicion.Equal(expected) {
		t.Errorf("acc.Phi(%d) != %#v, actual %#v", time, expected, suspicion)
	}
}
