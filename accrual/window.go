package accrual

import (
	"math"
)

type Window struct {
	elements []float64
	length   uint64
	head     uint64
}

func NewWindow(size uint64) Window {
	return Window{elements: make([]float64, size), head: 0}
}

func (w *Window) Push(n uint64) {
	if !w.full() {
		w.length++
	}

	w.elements[w.head] = float64(n)
	w.head = (w.head + 1) % uint64(len(w.elements))
}

func (w *Window) Len() uint64 {
	return w.length
}

func (w *Window) Mean() float64 {
	if w.Len() == 0 {
		return 0
	}

	var sum float64
	for _, v := range w.elements[:w.Len()] {
		sum += v
	}
	return sum / float64(w.Len())
}

func (w *Window) StdDev() float64 {
	mean := w.Mean()
	var ss float64
	for _, v := range w.elements[:w.Len()] {
		d := v - mean
		ss += d * d
	}
	return math.Sqrt(ss / float64(w.Len()))
}

func (w *Window) full() bool {
	return w.length == uint64(len(w.elements))
}
