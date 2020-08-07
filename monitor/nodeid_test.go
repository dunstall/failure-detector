package monitor

import (
	"reflect"
	"testing"
)

func TestNodeIDEncode(t *testing.T) {
	id := NodeID(0xaabbccdd)
	expected := []byte{0xaa, 0xbb, 0xcc, 0xdd}
	actual := id.Encode()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("id.Encode() != %v, actual %v", expected, actual)
	}
}

func TestNodeIDDecodeOk(t *testing.T) {
	encoded := []byte{0xaa, 0xbb, 0xcc, 0xdd}
	expected := NodeID(0xaabbccdd)

	var id NodeID
	if err := id.Decode(encoded); err != nil {
		t.Error(err)
	}

	if expected != id {
		t.Errorf("id.Decode(%v) != %#v, actual %#v", encoded, expected, id)
	}
}

func TestNodeIDDecodeErr(t *testing.T) {
	var id NodeID
	encodedTooSmall := []byte{1, 2}
	if err := id.Decode(encodedTooSmall); err != ErrInvalidID {
		t.Errorf("id.Decode(%v) != %v", encodedTooSmall, ErrInvalidID)
	}
}
