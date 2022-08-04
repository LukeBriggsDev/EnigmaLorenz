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

func TestRotorLengths(t *testing.T) {
	wheels := lorenz.NewWheelSet()
	lengths := []int{41, 31, 29, 26, 23, 61, 37, 43, 47, 51, 53, 59}
	for idx, wheel := range append(append(wheels.Chi[:], wheels.Motor[:]...), wheels.Psi[:]...) {
		if len(wheel.GetPins()) != lengths[idx] {
			t.Errorf("wheel %d has incorrect length of pins. Expected: %d, Actual: %d", idx, lengths[idx], len(wheel.GetPins()))
		}
	}
}

func TestLorenzEncrypt(t *testing.T) {
	wheels := lorenz.NewWheelSet()
	machine := lorenz.NewLorenz(
		wheels.Chi,
		wheels.Motor,
		wheels.Psi,
	)
	alphabet := lorenz.NewITA2LSB()
	plaintext := "ABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890A1B2C3"
	encoded, _ := alphabet.AsciiToITA2(plaintext)
	encrypted := machine.Encrypt(encoded)
	machine.ResetRotorPos()
	decrypted := machine.Encrypt(encrypted)
	decoded, _ := alphabet.ITA2ToAscii(decrypted)
	if plaintext != decoded {
		t.Errorf("%s != %s", plaintext, decrypted)
	}
}
