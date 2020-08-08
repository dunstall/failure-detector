package accrual

import (
	"math"
	"testing"
)

func TestAccrualPhiEmpty(t *testing.T) {
	acc := NewAccrual(5)
	var expectedPhi float64 = 0.0
	var expectedNSamples uint64 = 0
	phi, nSamples := acc.Phi(0)
	if phi != expectedPhi || nSamples != expectedNSamples {
		t.Errorf("acc.Phi(0) != %f, %d, actual %f, %d", expectedPhi, expectedNSamples, phi, nSamples)
	}
}

func TestAccrualPhiOneHeartbeat(t *testing.T) {
	acc := NewAccrual(5)
	acc.Heartbeat(100)
	var expectedPhi float64 = 0.0
	var expectedNSamples uint64 = 1
	phi, nSamples := acc.Phi(200)
	if phi != expectedPhi || nSamples != expectedNSamples {
		t.Errorf("acc.Phi(200) != %f, %d, actual %f, %d", expectedPhi, expectedNSamples, phi, nSamples)
	}
}

func TestAccrualPhiMultiHeartbeat(t *testing.T) {
	acc := NewAccrual(5)
	acc.Heartbeat(100)
	acc.Heartbeat(200)
	acc.Heartbeat(400)

	var expectedPhi float64 = 0.799546
	var expectedNSamples uint64 = 3
	phi, nSamples := acc.Phi(600)
	if math.Abs(phi-expectedPhi) > floatCompThreshold || nSamples != expectedNSamples {
		t.Errorf("acc.Phi(%d) != %f, %d, actual %f, %d", 600, expectedPhi, expectedNSamples, phi, nSamples)
	}

	expectedPhi = 11.892836
	phi, nSamples = acc.Phi(900)
	if math.Abs(phi-expectedPhi) > floatCompThreshold || nSamples != expectedNSamples {
		t.Errorf("acc.Phi(%d) != %f, %d, actual %f, %d", 900, expectedPhi, expectedNSamples, phi, nSamples)
	}

	acc.Heartbeat(1000)
	acc.Heartbeat(2000)
	acc.Heartbeat(2500)

	expectedPhi = 0.091436
	expectedNSamples = 6
	phi, nSamples = acc.Phi(2700)
	if math.Abs(phi-expectedPhi) > floatCompThreshold || nSamples != expectedNSamples {
		t.Errorf("acc.Phi(%d) != %f, %d, actual %f, %d", 2600, expectedPhi, expectedNSamples, phi, nSamples)
	}

	acc.Heartbeat(2701)
	acc.Heartbeat(2702)
	acc.Heartbeat(2703)

	expectedPhi = 0.262772
	phi, nSamples = acc.Phi(3000)
	if math.Abs(phi-expectedPhi) > floatCompThreshold || nSamples != expectedNSamples {
		t.Errorf("acc.Phi(%d) != %f, %d, actual %f, %d", 2600, expectedPhi, expectedNSamples, phi, nSamples)
	}
}
