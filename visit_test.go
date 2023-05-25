package main

import (
	"testing"
)

func TestNewVist(t *testing.T) {
	v := newVisit(1, "2023-05-27 14:00", 5, 16)

	if v.id != 1 {
		t.Errorf("Output %v not equal to expected '1'", v.id)
	}
}
