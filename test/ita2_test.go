package test

import (
	"bytes"
	"fmt"
	"testing"
)
import "EnigmaLorenz/pkg/lorenz"

func TestAsciiToITA2(t *testing.T) {
	alphabet := lorenz.NewITA2LSB()
	ita, _ := alphabet.AsciiToITA2("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789A1B2C3", false)
	expected := []byte{0x18, 0x13, 0x0e, 0x12, 0x10, 0x16, 0x0b, 0x05, 0x0c, 0x1a, 0x1e, 0x09, 0x07, 0x06, 0x03, 0x0d, 0x1d, 0x0a, 0x14, 0x01, 0x1c, 0x0f, 0x19, 0x17, 0x15, 0x11, 0x1b, 0x0d, 0x1d, 0x19, 0x10, 0x0a, 0x01, 0x15, 0x1c, 0x0c, 0x03, 0x1f, 0x18, 0x1b, 0x1d, 0x1f, 0x13, 0x1b, 0x19, 0x1f, 0x0e, 0x1b, 0x10}
	if bytes.Compare(ita, expected) != 0 {
		fmt.Printf("%x\n", ita)
		t.Errorf("%x != %s", ita, expected)
	}
}

func TestITA2ToAscii(t *testing.T) {
	alphabet := lorenz.NewITA2LSB()
	byteSequence := []byte{0x18, 0x13, 0x0e, 0x12, 0x10, 0x16, 0x0b, 0x05, 0x0c, 0x1a, 0x1e, 0x09, 0x07, 0x06, 0x03, 0x0d, 0x1d, 0x0a, 0x14, 0x01, 0x1c, 0x0f, 0x19, 0x17, 0x15, 0x11, 0x1b, 0x0d, 0x1d, 0x19, 0x10, 0x0a, 0x01, 0x15, 0x1c, 0x0c, 0x03, 0x1f, 0x18, 0x1b, 0x1d, 0x1f, 0x13, 0x1b, 0x19, 0x1f, 0x0e, 0x1b, 0x10}
	str, _ := alphabet.ITA2ToAscii(byteSequence, true)
	expected := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789A1B2C3"
	if str != expected {
		t.Errorf("%s != %s", str, expected)
	}
}

func TestReverse(t *testing.T) {
	alphabet := lorenz.NewITA2LSB()
	start := "*OKVIL"
	ita, _ := alphabet.AsciiToITA2(start, false)
	str, _ := alphabet.ITA2ToAscii(ita, false)
	if str != start {
		t.Errorf("%s != %s", start, str)
	}
}
