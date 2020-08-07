package monitor

import (
	"time"
)

type Heartbeat struct {
	NodeID   NodeID
	Received uint64
}

func NewHeartbeat(id NodeID) Heartbeat {
	return Heartbeat{NodeID: id, Received: uint64(time.Now().UnixNano() / 1000)}
}

type Monitor interface {
	Heartbeats() <-chan Heartbeat
	Close() error
}
