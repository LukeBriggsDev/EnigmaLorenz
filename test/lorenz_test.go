package test

import (
	"EnigmaLorenz/pkg/lorenz"
	"testing"
)

func TestWheelsToByte(t *testing.T) {
	chi1 := lorenz.NewWheel([]bool{true}, 0)
	chi2 := lorenz.NewWheel([]bool{false}, 0)
	chi3 := lorenz.NewWheel([]bool{false}, 0)
	chi4 := lorenz.NewWheel([]bool{true}, 0)
	chi5 := lorenz.NewWheel([]bool{true}, 0)
	expected := byte(19)
	if lorenz.WheelsToByte([]lorenz.Wheel{chi1, chi2, chi3, chi4, chi5}) != expected {
		t.Errorf("WheelsToByte does not return %d", expected)
	}
}
