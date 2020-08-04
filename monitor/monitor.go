package monitor

import (
	"time"
)

type Heartbeat struct {
	NodeID   string
	Received uint64
}

type Monitor interface {
	Heartbeats() chan<- Heartbeat
}
