package test

import (
	"bytes"
	"testing"
)
import "EnigmaLorenz/pkg/lorenz"

func TestAsciiToITA2(t *testing.T) {
	alphabet := lorenz.NewITA2()
	ita, _ := alphabet.AsciiToITA2("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789A1B2C3")
	expected := []byte{3, 25, 14, 9, 1, 13, 26, 20, 6, 11, 15, 18, 28, 12, 24, 22, 23, 10, 5, 16, 7, 30, 19, 29, 21, 17, 27, 22, 23, 19, 1, 10, 16, 21, 7, 6, 24, 31, 3, 27, 23, 31, 25, 27, 19, 31, 14, 27, 1}
	if bytes.Compare(ita, expected) != 0 {
		t.Errorf("%s != %s", ita, expected)
	}
}
