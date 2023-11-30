package main

import "testing"

func TestSum(t *testing.T) {
	got := sum(1, 1)
	if got != 2 {
		t.Errorf("Sum(1, 1) = %d; want 2", got)
	}
}
