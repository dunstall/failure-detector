package fd

import (
	"math"
	"testing"
)

const (
	floatCompThreshold = 1e-6
)

func TestWindowInitiallyLengthZero(t *testing.T) {
	w := NewWindow(4)
	var expectedLen uint64 = 0
	if w.Len() != expectedLen {
		t.Errorf("w.Len() != %d, actual %d", expectedLen, w.Len())
	}
}

func TestWindowEmpty(t *testing.T) {
	w := NewWindow(0)

	var expectedMean float64 = 0.0
	if w.Mean() != expectedMean {
		t.Errorf("w.Mean() != %f, actual %f", expectedMean, w.Mean())
	}

	var expectedStdDev float64 = 0
	if math.Abs(expectedStdDev-w.StdDev()) > floatCompThreshold {
		t.Errorf("w.StdDev() != %f, actual %f", expectedStdDev, w.StdDev())
	}
}

func TestWindowOneElement(t *testing.T) {
	w := NewWindow(4)
	w.Push(5)

	var expectedLen uint64 = 1
	if w.Len() != expectedLen {
		t.Errorf("w.Len() != %d, actual %d", expectedLen, w.Len())
	}

	var expectedMean float64 = 5.0
	if w.Mean() != expectedMean {
		t.Errorf("w.Mean() != %f, actual %f", expectedMean, w.Mean())
	}

	var expectedStdDev float64 = 0
	if math.Abs(expectedStdDev-w.StdDev()) > floatCompThreshold {
		t.Errorf("w.StdDev() != %f, actual %f", expectedStdDev, w.StdDev())
	}
}

func TestWindowLenElements(t *testing.T) {
	w := NewWindow(4)
	w.Push(1)
	w.Push(2)
	w.Push(3)
	w.Push(4)

	var expectedLen uint64 = 4
	if w.Len() != expectedLen {
		t.Errorf("w.Len() != %d, actual %d", expectedLen, w.Len())
	}

	var expectedMean float64 = 2.5
	if w.Mean() != expectedMean {
		t.Errorf("w.Mean() != %f, actual %f", expectedMean, w.Mean())
	}

	var expectedStdDev float64 = 1.118034
	if math.Abs(expectedStdDev-w.StdDev()) > floatCompThreshold {
		t.Errorf("w.StdDev() != %f, actual %f", expectedStdDev, w.StdDev())
	}
}

func TestWindowExceedLenElements(t *testing.T) {
	w := NewWindow(2)
	w.Push(1)
	w.Push(2)
	w.Push(11)
	w.Push(12)

	var expectedLen uint64 = 2
	if w.Len() != expectedLen {
		t.Errorf("w.Len() != %d, actual %d", expectedLen, w.Len())
	}

	var expectedMean float64 = 11.5
	if w.Mean() != expectedMean {
		t.Errorf("w.Mean() != %f, actual %f", expectedMean, w.Mean())
	}

	var expectedStdDev float64 = 0.5
	if w.StdDev() != expectedStdDev {
		t.Errorf("w.StdDev() != %f, actual %f", expectedStdDev, w.StdDev())
	}
}
