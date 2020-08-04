package detector

import (
	"math"
	"sync"
)

type Detector struct {
	window Window
	last   uint64
	mu     sync.Mutex
}

func NewDetector(size uint64) Detector {
	return Detector{window: NewWindow(size)}
}

func (d *Detector) Heartbeat(t uint64) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.last != 0 {
		d.window.Push(t - d.last)
	}
	d.last = t
}

// Phi returns the suspicion level of the failure detector and the number of
// samples used.
func (d *Detector) Phi(t uint64) (float64, uint64) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.window.Len() == 0 {
		if d.last == 0 {
			return 0, 0
		} else {
			return 0, 1
		}
	}

	diff := t - d.last
	pLater := 1 - cdf(d.window.Mean(), d.window.StdDev(), float64(diff))
	phi := -math.Log10(pLater)

	return phi, d.window.Len() + 1
}

func cdf(mean, stddev, x float64) float64 {
	return 0.5 + 0.5*math.Erf((x-mean)/(stddev*math.Sqrt2))
}
