package main

import "testing"

func TestSquare(t *testing.T) {
	rst := square(9)
	if rst != 81 {
		t.Errorf("Square(9) should be 81 but square(9) returns %d", rst)
	}
}

func TestSquare23(t *testing.T) {
	rst := square(3)
	if rst != 9 {
		t.Errorf("Square(3) should be 9 but square(3) returns %d", rst)
	}
}
