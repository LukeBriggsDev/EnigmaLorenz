package enigma

import (
	"errors"
	"log"
)

func indexOf(needle byte, haystack []byte) (byte, error) {
	for idx, val := range haystack {
		if val == needle {
			return byte(idx), nil
		}
	}
	return 0, errors.New("Item not found in list")
}

func negMod(a int, n int) int {
	return (a%n + n) % n
}

type Rotor struct {
	Name         string
	Wires        []byte
	shownPos     byte
	TurnoverList []byte
	ringSetting  byte
}

func (r *Rotor) SetShownPos(pos byte) {
	if pos < 1 || pos > 26 {
		log.Fatal("Rotor position must be set to value between 1 and 26")
	}
	r.shownPos = pos - 1
}

func (r *Rotor) GetShownPos() byte {
	return r.shownPos + 1
}

func (r *Rotor) GetRingSetting() byte {
	return r.ringSetting
}

func (r *Rotor) SetRingSetting(offset byte) {
	if offset < 0 || offset > 25 {
		log.Fatal("Rotor position must be set to value between 0 and 25")
	}
	r.ringSetting = offset
}

func (r *Rotor) NormalizedPos() byte {
	return byte(negMod(int(r.shownPos)-int(r.ringSetting), len(r.Wires)))
}

func (r *Rotor) AtNotch() bool {
	_, err := indexOf(r.shownPos, r.TurnoverList)
	return err == nil
}

func (r *Rotor) Rotate() {
	r.shownPos += 1
	r.shownPos %= byte(len(r.Wires))
}

func (r Rotor) Translate(plain byte) byte {

	shiftedPlain := (plain + r.NormalizedPos()) % byte(len(r.Wires)) // Shift by rotor and ring position
	shiftedCipher := negMod(int(r.Wires[shiftedPlain])-int(r.NormalizedPos()), len(r.Wires))
	return byte(shiftedCipher)
}

func (r Rotor) TranslateReverse(cipher byte) byte {
	unshiftedCipher := (cipher + r.NormalizedPos()) % byte(len(r.Wires))
	unshiftedPlain, _ := indexOf(unshiftedCipher, r.Wires)
	shiftedPlain := negMod(int(unshiftedPlain)-int(r.NormalizedPos()), len(r.Wires))
	return byte(shiftedPlain)
}

func GenerateRotors() (I Rotor, II Rotor, III Rotor, UKW_A Rotor, UKW_B Rotor) {
	I = Rotor{
		Name:         "I",
		Wires:        []byte{4, 10, 12, 5, 11, 6, 3, 16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9},
		shownPos:     0,
		TurnoverList: []byte{16},
	}

	II = Rotor{
		Name:         "II",
		Wires:        []byte{0, 9, 3, 10, 18, 8, 17, 20, 23, 1, 11, 7, 22, 19, 12, 2, 16, 6, 25, 13, 15, 24, 5, 21, 14, 4},
		shownPos:     0,
		TurnoverList: []byte{4},
	}

	III = Rotor{
		Name:         "III",
		Wires:        []byte{1, 3, 5, 7, 9, 11, 2, 15, 17, 19, 23, 21, 25, 13, 24, 4, 8, 22, 6, 0, 10, 12, 20, 18, 16, 14},
		shownPos:     0,
		TurnoverList: []byte{21},
	}

	UKW_A = Rotor{
		Name:  "UKW-A",
		Wires: []byte{4, 9, 12, 25, 0, 11, 24, 23, 21, 1, 22, 5, 2, 17, 16, 20, 14, 13, 19, 18, 15, 8, 10, 7, 6, 3},
	}

	UKW_B = Rotor{
		Name:  "UKW-B",
		Wires: []byte{24, 17, 20, 7, 16, 18, 11, 3, 15, 23, 13, 6, 14, 10, 12, 8, 4, 1, 5, 25, 2, 22, 21, 9, 0, 19},
	}
	return
}
