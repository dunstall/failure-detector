package accrual

import (
	"math"
	"sync"
)

const (
	floatCompThreshold = 1e-6
)

type Suspicion struct {
	Phi      float64
	NSamples uint64
}

func (s *Suspicion) Equal(u Suspicion) bool {
	return math.Abs(u.Phi-s.Phi) < floatCompThreshold && u.NSamples == s.NSamples
}

type Accrual struct {
	window Window
	last   uint64
	mu     sync.Mutex // TODO RWMutex
}

func NewAccrual(size uint64) Accrual {
	return Accrual{window: NewWindow(size)}
}

func (acc *Accrual) Heartbeat(t uint64) {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	if acc.last != 0 {
		acc.window.Push(t - acc.last)
	}
	acc.last = t
}

// Phi returns the suspicion level of the failure detector and the number of
// samples used.
func (acc *Accrual) Phi(t uint64) Suspicion {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	if acc.window.Len() == 0 {
		if acc.last == 0 {
			return Suspicion{0, 0}
		} else {
			return Suspicion{0, 1}
		}
	}

	if t < acc.last {
		return Suspicion{0, acc.window.Len() + 1}
	}

	// Adapted from https://github.com/dgryski/go-failure.
	diff := t - acc.last
	pLater := 1 - cdf(acc.window.Mean(), acc.window.StdDev(), float64(diff))
	phi := -math.Log10(pLater)

	return Suspicion{phi, acc.window.Len() + 1}
}

func cdf(mean, stddev, x float64) float64 {
	return 0.5 + 0.5*math.Erf((x-mean)/(stddev*math.Sqrt2))
}
