package main

import (
	"testing"
	"time"
)

func TestDefaultTimeToUnixNanoAndBackAgain(t *testing.T) {
	var defaultTime time.Time // 0001-01-01 00:00:00 +0000 UTC

	encodedDefaultTime := defaultTime.UnixNano() //gives -6795364578871345152

	decodedDefaultTime := time.Unix(0, encodedDefaultTime)

	if defaultTime != decodedDefaultTime {
		t.Errorf("expected: %v, got %v", defaultTime, decodedDefaultTime)
	}
}

func TestDefaultTimeToUnixAndBackAgain(t *testing.T) {
	var defaultTime time.Time // 0001-01-01 00:00:00 +0000 UTC

	encodedDefaultTime := defaultTime.Unix() //gives -62135596800

	decodedDefaultTime := time.Unix(encodedDefaultTime, 0)

	if defaultTime != decodedDefaultTime {
		t.Errorf("val:expected: %v, got %v", defaultTime, decodedDefaultTime)
	}

	if defaultTime.Unix() != decodedDefaultTime.Unix() {
		t.Errorf("Unix():expected: %v, got %v", defaultTime.Unix(), decodedDefaultTime.Unix())
	}

}
