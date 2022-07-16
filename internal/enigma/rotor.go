package enigma

import "errors"

func indexOf(needle byte, haystack []byte) (byte, error) {
	for idx, val := range haystack {
		if val == needle {
			return byte(idx), nil
		}
	}
	return 0, errors.New("Item not found in list")
}

type Rotor struct {
	Name         string
	Alphabet     []byte
	CurrentPos   byte
	NotchList    []byte
	TurnoverList []byte
}

func (r *Rotor) Rotate() {
	r.CurrentPos += 1
}

func (r Rotor) Translate(plain byte) byte {

	shiftedPlain := (plain + r.CurrentPos) % byte(len(r.Alphabet))
	shiftedCipher := (r.Alphabet[shiftedPlain] - r.CurrentPos + byte(len(r.Alphabet))) % byte(len(r.Alphabet))
	return shiftedCipher
}

func (r Rotor) TranslateReverse(cipher byte) byte {
	unshiftedCipher := (cipher + r.CurrentPos) % byte(len(r.Alphabet))
	unshiftedPlain, _ := indexOf(unshiftedCipher, r.Alphabet)
	shiftedPlain := (unshiftedPlain - r.CurrentPos + byte(len(r.Alphabet))) % byte(len(r.Alphabet))
	return shiftedPlain
}

func GenerateRotors() (I Rotor, II Rotor, III Rotor, UKW_A Rotor, UKW_B Rotor) {
	I = Rotor{
		Name:         "I",
		Alphabet:     []byte{4, 10, 12, 5, 11, 6, 3, 16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9},
		CurrentPos:   0,
		NotchList:    []byte{'Y'},
		TurnoverList: []byte{'Q'},
	}

	II = Rotor{
		Name:         "II",
		Alphabet:     []byte{0, 9, 3, 10, 18, 8, 17, 20, 23, 1, 11, 7, 22, 19, 12, 2, 16, 6, 25, 13, 15, 24, 5, 21, 14, 4},
		CurrentPos:   0,
		NotchList:    []byte{'M'},
		TurnoverList: []byte{'E'},
	}

	III = Rotor{
		Name:         "III",
		Alphabet:     []byte{1, 3, 5, 7, 9, 11, 2, 15, 17, 19, 23, 21, 25, 13, 24, 4, 8, 22, 6, 0, 10, 12, 20, 18, 16, 14},
		CurrentPos:   0,
		NotchList:    []byte{'D'},
		TurnoverList: []byte{'V'},
	}

	UKW_A = Rotor{
		Name:     "UKW-A",
		Alphabet: []byte{4, 9, 12, 25, 0, 11, 24, 23, 21, 1, 22, 5, 2, 17, 16, 20, 14, 13, 19, 18, 15, 8, 10, 7, 6, 3},
	}

	UKW_B = Rotor{
		Name:     "UKW-B",
		Alphabet: []byte{24, 17, 20, 7, 16, 18, 11, 3, 15, 23, 13, 6, 14, 10, 12, 8, 4, 1, 5, 25, 2, 22, 21, 9, 0, 19},
	}
	return
}
