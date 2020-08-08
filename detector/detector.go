package detector

import (
	"errors"
	"sync"

	"github.com/dunstall/failure-detector/accrual"
	"github.com/dunstall/failure-detector/monitor"
)

var (
	ErrNotFound = errors.New("node not found (never received heartbeat)")
)

type Detector struct {
	monitor     monitor.Monitor
	accrual     map[monitor.NodeID]*accrual.Accrual
	done        chan bool
	mu          sync.RWMutex
	accrualSize uint64
}

func NewDetector(m monitor.Monitor, accrualSize uint64) Detector {
	d := Detector{
		monitor:     m,
		accrual:     make(map[monitor.NodeID]*accrual.Accrual),
		done:        make(chan bool),
		accrualSize: accrualSize,
	}
	go d.run()
	return d
}

func (d *Detector) Phi(t uint64, id monitor.NodeID) (accrual.Suspicion, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if acc, ok := d.accrual[id]; ok {
		return acc.Phi(t), nil
	}
	return accrual.Suspicion{}, ErrNotFound
}

func (d *Detector) Close() {
	close(d.done)
	d.monitor.Close()
}

func (d *Detector) run() {
	for {
		select {
		case <-d.done:
			return
		case hb := <-d.monitor.Heartbeats():
			d.heartbeat(hb)
		}
	}
}

func (d *Detector) heartbeat(hb monitor.Heartbeat) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, ok := d.accrual[hb.NodeID]; !ok {
		acc := accrual.NewAccrual(d.accrualSize)
		d.accrual[hb.NodeID] = &acc
	}
	d.accrual[hb.NodeID].Heartbeat(hb.Received)
}
