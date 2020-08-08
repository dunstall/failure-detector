package accrual

import (
	// "math"
	"testing"
)

func TestAccrualPhiEmpty(t *testing.T) {
	acc := NewAccrual(5)

	expected := Suspicion{0.0, 0}

	var time uint64 = 0
	suspicion := acc.Phi(time)
	if !suspicion.Equal(expected) {
		t.Errorf("acc.Phi(%d) != %#v, actual %#v", time, expected, suspicion)
	}
}

func TestAccrualPhiOneHeartbeat(t *testing.T) {
	acc := NewAccrual(5)
	acc.Heartbeat(100)

	expected := Suspicion{0.0, 1}

	var time uint64 = 200
	suspicion := acc.Phi(time)
	if !suspicion.Equal(expected) {
		t.Errorf("acc.Phi(%d) != %#v, actual %#v", time, expected, suspicion)
	}
}

func TestAccrualPhiBeforeLast(t *testing.T) {
	acc := NewAccrual(5)
	acc.Heartbeat(100)
	acc.Heartbeat(200)

	expected := Suspicion{0.0, 2}

	var time uint64 = 150
	suspicion := acc.Phi(time)
	if !suspicion.Equal(expected) {
		t.Errorf("acc.Phi(%d) != %#v, actual %#v", time, expected, suspicion)
	}
}

func TestAccrualPhiMultiHeartbeat(t *testing.T) {
	acc := NewAccrual(5)
	acc.Heartbeat(100)
	acc.Heartbeat(200)
	acc.Heartbeat(400)

	expected := Suspicion{0.799546, 3}
	var time uint64 = 600
	suspicion := acc.Phi(time)
	if !suspicion.Equal(expected) {
		t.Errorf("acc.Phi(%d) != %#v, actual %#v", time, expected, suspicion)
	}

	expected = Suspicion{11.892836, 3}
	time = 900
	suspicion = acc.Phi(time)
	if !suspicion.Equal(expected) {
		t.Errorf("acc.Phi(%d) != %#v, actual %#v", time, expected, suspicion)
	}

	acc.Heartbeat(1000)
	acc.Heartbeat(2000)
	acc.Heartbeat(2500)

	expected = Suspicion{0.091436, 6}
	time = 2700
	suspicion = acc.Phi(time)
	if !suspicion.Equal(expected) {
		t.Errorf("acc.Phi(%d) != %#v, actual %#v", time, expected, suspicion)
	}

	acc.Heartbeat(2701)
	acc.Heartbeat(2702)
	acc.Heartbeat(2703)

	expected = Suspicion{0.262772, 6}
	time = 3000
	suspicion = acc.Phi(time)
	if !suspicion.Equal(expected) {
		t.Errorf("acc.Phi(%d) != %#v, actual %#v", time, expected, suspicion)
	}
}
