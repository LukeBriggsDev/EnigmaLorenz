package test

import "testing"
import "EnigmaLorenz/internal/enigma"

func TestRotorTranslateNoOffset(t *testing.T) {
	I, _, _, _, _ := enigma.GenerateRotors()
	cipher := I.Translate(0)
	if cipher != 4 {
		t.Errorf("Rotor %s. Expected: 0 -> 4. Got: A -> %d", I.Name, cipher)
	}

}

func TestRotorTranslateReverseNoOffset(t *testing.T) {
	I, _, _, _, _ := enigma.GenerateRotors()
	cipher := I.Translate(0)
	plain := I.TranslateReverse(cipher)
	if plain != 0 {
		t.Errorf("Incorrect reverse translation for rotor %s, %d comes back as %d", I.Name, 0, plain)
	}
}

func TestRotorTranslateReverseOffset(t *testing.T) {
	I, _, _, _, _ := enigma.GenerateRotors()
	for chr := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		for i := 0; i < 26; i++ {
			I.CurrentPos = byte(i)
			cipher := I.Translate(byte(chr))
			plain := I.TranslateReverse(cipher)
			if byte(chr) != plain {
				t.Errorf("Rotor %s Offset %d. Expected: %d -> %d. Got: %d -> %d", I.Name, i, chr, chr, chr, plain)
			}
		}
	}
}
