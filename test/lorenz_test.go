package test

import (
	"EnigmaLorenz/pkg/lorenz"
	"fmt"
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

func TestRotorLengths(t *testing.T) {
	wheels := lorenz.NewWheelSet()
	lengths := []int{41, 31, 29, 26, 23, 61, 37, 43, 47, 51, 53, 59}
	for idx, wheel := range []lorenz.Wheel{wheels.Chi1, wheels.Chi2, wheels.Chi3, wheels.Chi4, wheels.Chi5, wheels.M1, wheels.M2, wheels.Psi1, wheels.Psi2, wheels.Psi3, wheels.Psi4, wheels.Psi5} {
		if len(wheel.GetPins()) != lengths[idx] {
			t.Errorf("wheel %d has incorrect length of pins. Expected: %d, Actual: %d", idx, lengths[idx], len(wheel.GetPins()))
		}
	}
}

func TestLorenzEncrypt(t *testing.T) {
	wheels := lorenz.NewWheelSet()
	machine := lorenz.NewLorenz(
		[]lorenz.Wheel{wheels.Chi1, wheels.Chi2, wheels.Chi3, wheels.Chi4, wheels.Chi5},
		wheels.M1,
		wheels.M2,
		[]lorenz.Wheel{wheels.Psi1, wheels.Psi2, wheels.Psi3, wheels.Psi4, wheels.Psi5},
	)
	alphabet := lorenz.NewITA2MSB()
	encoded, _ := alphabet.AsciiToITA2("ABC")
	encrypted := machine.Encrypt(encoded)
	decoded, _ := alphabet.ITA2ToAscii(encrypted)
	fmt.Println(decoded)
}
