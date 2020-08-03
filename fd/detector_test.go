package fd

import (
	"math"
	"testing"
)

func TestDetectorPhiEmpty(t *testing.T) {
	d := NewDetector(5)
	var expectedPhi float64 = 0.0
	var expectedNSamples uint64 = 0
	phi, nSamples := d.Phi(0)
	if phi != expectedPhi || nSamples != expectedNSamples {
		t.Errorf("w.Len() != %f, %d, actual %f, %d", expectedPhi, expectedNSamples, phi, nSamples)
	}
}

func TestDetectorPhiOneHeartbeat(t *testing.T) {
	d := NewDetector(5)
	d.Heartbeat(100)
	var expectedPhi float64 = 0.0
	var expectedNSamples uint64 = 1
	phi, nSamples := d.Phi(200)
	if phi != expectedPhi || nSamples != expectedNSamples {
		t.Errorf("w.Len() != %f, %d, actual %f, %d", expectedPhi, expectedNSamples, phi, nSamples)
	}
}

func TestDetectorPhiMultiHeartbeat(t *testing.T) {
	d := NewDetector(5)
	d.Heartbeat(100)
	d.Heartbeat(200)
	d.Heartbeat(400)

	var expectedPhi float64 = 0.799546
	var expectedNSamples uint64 = 3
	phi, nSamples := d.Phi(600)
	if math.Abs(phi-expectedPhi) > floatCompThreshold || nSamples != expectedNSamples {
		t.Errorf("d.Phi(%d) != %f, %d, actual %f, %d", 600, expectedPhi, expectedNSamples, phi, nSamples)
	}

	expectedPhi = 11.892836
	phi, nSamples = d.Phi(900)
	if math.Abs(phi-expectedPhi) > floatCompThreshold || nSamples != expectedNSamples {
		t.Errorf("w.Phi(%d) != %f, %d, actual %f, %d", 900, expectedPhi, expectedNSamples, phi, nSamples)
	}

	d.Heartbeat(1000)
	d.Heartbeat(2000)
	d.Heartbeat(2500)

	expectedPhi = 0.091436
	expectedNSamples = 6
	phi, nSamples = d.Phi(2700)
	if math.Abs(phi-expectedPhi) > floatCompThreshold || nSamples != expectedNSamples {
		t.Errorf("w.Phi(%d) != %f, %d, actual %f, %d", 2600, expectedPhi, expectedNSamples, phi, nSamples)
	}

	d.Heartbeat(2701)
	d.Heartbeat(2702)
	d.Heartbeat(2703)

	expectedPhi = 0.262772
	phi, nSamples = d.Phi(3000)
	if math.Abs(phi-expectedPhi) > floatCompThreshold || nSamples != expectedNSamples {
		t.Errorf("w.Phi(%d) != %f, %d, actual %f, %d", 2600, expectedPhi, expectedNSamples, phi, nSamples)
	}
}
