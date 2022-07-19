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
			I.ShownPos = byte(i)
			cipher := I.Translate(byte(chr))
			plain := I.TranslateReverse(cipher)
			if byte(chr) != plain {
				t.Errorf("Rotor %s Offset %d. Expected: %d -> %d. Got: %d -> %d", I.Name, i, chr, chr, chr, plain)
			}
		}
	}
}

func TestMachineEncrypt(t *testing.T) {
	I, II, III, _, UKW_B := enigma.GenerateRotors()
	machine := enigma.Enigma{
		LeftRotor:   III,
		CenterRotor: II,
		RightRotor:  I,
		Reflector:   UKW_B,
	}
	plaintext := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	expectedCipher := "FTZMGISXIPJWGDNJJCOQTYRIGDMXFIESRWZGTOIUIEKKDCSHTPYOEPVXNHVRWWESFRUXDGWOZDMNKIZWNCZDUCOBLTUYHDZGO"
	cipher := machine.Encrypt(plaintext)
	if cipher != expectedCipher {
		t.Errorf("Machine: %s, %s, %s, %s.\nPlaintext:\t\t\t%s.\nExpected Cipher:\t%s.\nActual Cipher:\t\t%s.\n", UKW_B.Name, III.Name, II.Name, I.Name, plaintext, expectedCipher, cipher)
	}
}

func TestMachineEncryptRing(t *testing.T) {
	I, II, III, _, UKW_B := enigma.GenerateRotors()
	I.RingSetting = 1
	II.RingSetting = 3
	III.RingSetting = 6
	machine := enigma.Enigma{
		LeftRotor:   III,
		CenterRotor: II,
		RightRotor:  I,
		Reflector:   UKW_B,
	}
	plaintext := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	expectedCipher := "MBPEKFULTQTRXBRUSTTUDGKWSTJSGJVYWBIGVUEJBKHOLKENMWVUXIQJZXQWNCZMERMKRMRDGYQEREZCTPWTJXQLEEKKDCZXX"
	cipher := machine.Encrypt(plaintext)
	if cipher != expectedCipher {
		t.Errorf("Machine: %s, %s, %s, %s.\nPlaintext:\t\t\t%s.\nExpected Cipher:\t%s.\nActual Cipher:\t\t%s.\n", UKW_B.Name, III.Name, II.Name, I.Name, plaintext, expectedCipher, cipher)
	}
}

func TestPlugboard(t *testing.T) {
	plugboard := enigma.NewPlugboard()
	plugboard.AddPlug('A', 'B')
	result := plugboard.Translate('A')
	if result != 'B' {
		t.Errorf("A should translate to B, instead A becomes %c", result)
	}
	plugboard.RemovePlug('A')
	result = plugboard.Translate('A')
	if result != 'A' {
		t.Errorf("A should become A with empty plugboard, instead A becomes %c", result)
	}
}
