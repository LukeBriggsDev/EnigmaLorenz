package test

import "testing"
import "EnigmaLorenz/internal/enigma"

func TestRotorTranslateNoOffset(t *testing.T) {
	rotorSet := enigma.GenerateRotors()
	I := rotorSet.I
	cipher := I.Translate(0)
	if cipher != 4 {
		t.Errorf("Rotor %s. Expected: 0 -> 4. Got: A -> %d", I.Name, cipher)
	}

}

func TestRotorTranslateReverseNoOffset(t *testing.T) {
	rotorSet := enigma.GenerateRotors()
	I := rotorSet.I
	cipher := I.Translate(0)
	plain := I.TranslateReverse(cipher)
	if plain != 0 {
		t.Errorf("Incorrect reverse translation for rotor %s, %d comes back as %d", I.Name, 0, plain)
	}
}

func TestRotorTranslateReverseOffset(t *testing.T) {
	rotorSet := enigma.GenerateRotors()
	I := rotorSet.I
	for chr := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		for i := 1; i < 27; i++ {
			I.SetShownPos(byte(i))
			cipher := I.Translate(byte(chr))
			plain := I.TranslateReverse(cipher)
			if byte(chr) != plain {
				t.Errorf("Rotor %s Offset %d. Expected: %d -> %d. Got: %d -> %d", I.Name, i, chr, chr, chr, plain)
			}
		}
	}
}

func TestMachineEncrypt(t *testing.T) {
	rotorSet := enigma.GenerateRotors()
	I := rotorSet.I
	II := rotorSet.II
	III := rotorSet.III
	UKW_B := rotorSet.UKW_B
	machine := enigma.Enigma{
		LeftRotor:   III,
		CenterRotor: II,
		RightRotor:  I,
		Reflector:   UKW_B,
	}
	plaintext := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	expectedCipher := "FTZMGISXIPJWGDNJJCOQTYRIGDMXFIESRWZGTOIUIEKKDCSHTPYOEPVXNHVRWWESFRUXDGWOZDMNKIZWNCZDUCOBLTUYHDZGO"
	cipher, _ := machine.Encrypt(plaintext, false)
	if cipher != expectedCipher {
		t.Errorf("Machine: %s, %s, %s, %s.\nPlaintext:\t\t\t%s.\nExpected Cipher:\t%s.\nActual Cipher:\t\t%s.\n", UKW_B.Name, III.Name, II.Name, I.Name, plaintext, expectedCipher, cipher)
	}
}

func TestMachineEncryptFourRotor(t *testing.T) {
	rotorSet := enigma.GenerateRotors()
	UKW_B := rotorSet.UKW_B
	machine := enigma.Enigma{
		LeftRotor:   rotorSet.III,
		CenterRotor: rotorSet.II,
		RightRotor:  rotorSet.I,
		FourthRotor: rotorSet.Beta,
		Reflector:   rotorSet.UKW_b,
	}
	plaintext := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	expectedCipher := "FTZMGISXIPJWGDNJJCOQTYRIGDMXFIESRWZGTOIUIEKKDCSHTPYOEPVXNHVRWWESFRUXDGWOZDMNKIZWNCZDUCOBLTUYHDZGO"
	cipher, _ := machine.Encrypt(plaintext, true)
	if cipher != expectedCipher {
		t.Errorf("Machine: %s, %s, %s, %s.\nPlaintext:\t\t\t%s.\nExpected Cipher:\t%s.\nActual Cipher:\t\t%s.\n", UKW_B.Name, rotorSet.III.Name, rotorSet.II.Name, rotorSet.I.Name, plaintext, expectedCipher, cipher)
	}
}

func TestMachineEncryptRing(t *testing.T) {
	rotorSet := enigma.GenerateRotors()
	I := rotorSet.I
	II := rotorSet.II
	III := rotorSet.III
	UKW_B := rotorSet.UKW_B
	I.SetRingSetting(1)
	II.SetRingSetting(3)
	III.SetRingSetting(6)
	machine := enigma.Enigma{
		LeftRotor:   III,
		CenterRotor: II,
		RightRotor:  I,
		Reflector:   UKW_B,
	}
	plaintext := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	expectedCipher := "MBPEKFULTQTRXBRUSTTUDGKWSTJSGJVYWBIGVUEJBKHOLKENMWVUXIQJZXQWNCZMERMKRMRDGYQEREZCTPWTJXQLEEKKDCZXX"
	cipher, _ := machine.Encrypt(plaintext, false)
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

func TestMachineEncryptRingPlugboard(t *testing.T) {
	rotorSet := enigma.GenerateRotors()
	I := rotorSet.I
	II := rotorSet.II
	III := rotorSet.III
	UKW_B := rotorSet.UKW_B
	I.SetRingSetting(6)
	II.SetRingSetting(3)
	III.SetRingSetting(1)
	machine := enigma.Enigma{
		LeftRotor:   III,
		CenterRotor: II,
		RightRotor:  I,
		Reflector:   UKW_B,
		Plugs:       enigma.NewPlugboard(),
	}
	plug := "STEADING"
	for i := 0; i < len(plug); i += 2 {
		machine.Plugs.AddPlug(plug[i], plug[i+1])
	}
	plaintext := "QWERTYUIOPASDFGHJKLZXCVBNMMNBVCXZLKJHGFDSAPOIUYTREWQQWERTYUIOPASDFGHJKLZXCVBNMMNBVCXZLKJHGFDSAPOIUYTREWQ"
	expectedCipher := "ALOMPLZEIIRAHXECCABOCJAHYUAUQKGQJXEQUVPIZSBLWPQFOPEWGFUTAWRVYSSIWTIIWLFKAUUZYWYUFPIOEOHEQSCCJDXBQISSCKSW"
	cipher, _ := machine.Encrypt(plaintext, false)
	if cipher != expectedCipher {
		t.Errorf("Machine: %s, %s, %s, %s.\nPlaintext:\t\t\t%s.\nExpected Cipher:\t%s.\nActual Cipher:\t\t%s.\n", UKW_B.Name, III.Name, II.Name, I.Name, plaintext, expectedCipher, cipher)
	}
}
