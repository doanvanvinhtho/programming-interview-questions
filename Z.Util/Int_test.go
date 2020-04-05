package util

import "testing"

func TestIntMax(t *testing.T) {
	if IntMax(1, 1) != 1 {
		t.Errorf("Wrong 1")
	}
	if IntMax(1, 2) != 2 {
		t.Errorf("Wrong 2")
	}
}
