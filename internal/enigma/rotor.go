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

type RotorSet struct {
	I     Rotor
	II    Rotor
	III   Rotor
	IV    Rotor
	V     Rotor
	VI    Rotor
	VII   Rotor
	VIII  Rotor
	UKW_A Rotor
	UKW_B Rotor
	UKW_C Rotor
	UKW_b Rotor
	UKW_c Rotor
	Beta  Rotor
	Gamma Rotor
}

func GenerateRotors() (set RotorSet) {
	I := Rotor{
		Name:         "I",
		Wires:        []byte{4, 10, 12, 5, 11, 6, 3, 16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9},
		shownPos:     0,
		TurnoverList: []byte{16},
	}

	II := Rotor{
		Name:         "II",
		Wires:        []byte{0, 9, 3, 10, 18, 8, 17, 20, 23, 1, 11, 7, 22, 19, 12, 2, 16, 6, 25, 13, 15, 24, 5, 21, 14, 4},
		shownPos:     0,
		TurnoverList: []byte{4},
	}

	III := Rotor{
		Name:         "III",
		Wires:        []byte{1, 3, 5, 7, 9, 11, 2, 15, 17, 19, 23, 21, 25, 13, 24, 4, 8, 22, 6, 0, 10, 12, 20, 18, 16, 14},
		shownPos:     0,
		TurnoverList: []byte{21},
	}

	IV := Rotor{
		Name:         "IV",
		Wires:        []byte{4, 18, 14, 21, 15, 25, 9, 0, 24, 16, 20, 8, 17, 7, 23, 11, 13, 5, 19, 6, 10, 3, 2, 12, 22, 1},
		shownPos:     0,
		TurnoverList: []byte{9},
	}

	V := Rotor{
		Name:         "V",
		Wires:        []byte{21, 25, 1, 17, 6, 8, 19, 24, 20, 15, 18, 3, 13, 7, 11, 23, 0, 22, 12, 9, 16, 14, 5, 4, 2, 10},
		shownPos:     0,
		TurnoverList: []byte{25},
	}

	VI := Rotor{
		Name:         "VI",
		Wires:        []byte{9, 15, 6, 21, 14, 20, 12, 5, 24, 16, 1, 4, 13, 7, 25, 17, 3, 10, 0, 18, 23, 11, 8, 2, 19, 22},
		shownPos:     0,
		TurnoverList: []byte{25, 7},
	}

	VII := Rotor{
		Name:         "VII",
		Wires:        []byte{13, 25, 9, 7, 6, 17, 2, 23, 12, 24, 18, 22, 1, 14, 20, 5, 0, 8, 21, 11, 15, 4, 10, 16, 3, 19},
		shownPos:     0,
		TurnoverList: []byte{25, 7},
	}

	VIII := Rotor{
		Name:         "VIII",
		Wires:        []byte{5, 10, 16, 7, 19, 11, 23, 14, 2, 1, 9, 18, 15, 3, 25, 17, 0, 12, 4, 22, 13, 8, 20, 24, 6, 21},
		shownPos:     0,
		TurnoverList: []byte{25, 7},
	}

	UKW_A := Rotor{
		Name:  "UKW-A",
		Wires: []byte{4, 9, 12, 25, 0, 11, 24, 23, 21, 1, 22, 5, 2, 17, 16, 20, 14, 13, 19, 18, 15, 8, 10, 7, 6, 3},
	}

	UKW_B := Rotor{
		Name:  "UKW-B",
		Wires: []byte{24, 17, 20, 7, 16, 18, 11, 3, 15, 23, 13, 6, 14, 10, 12, 8, 4, 1, 5, 25, 2, 22, 21, 9, 0, 19},
	}

	UKW_C := Rotor{
		Name:  "UKW-C",
		Wires: []byte{17, 3, 14, 1, 9, 13, 19, 10, 21, 4, 7, 12, 11, 5, 2, 22, 25, 0, 23, 6, 24, 8, 15, 18, 20, 16},
	}

	Beta := Rotor{
		Name:  "Beta",
		Wires: []byte{11, 4, 24, 9, 21, 2, 13, 8, 23, 22, 15, 1, 16, 12, 3, 17, 19, 0, 10, 25, 6, 5, 20, 7, 14, 18},
	}

	Gamma := Rotor{
		Name:  "Gamma",
		Wires: []byte{5, 18, 14, 10, 0, 13, 20, 4, 17, 7, 12, 1, 19, 8, 24, 2, 22, 11, 16, 15, 25, 23, 21, 6, 9, 3},
	}

	UKW_b := Rotor{
		Name:  "Narrow UKW-b",
		Wires: []byte{4, 13, 10, 16, 0, 20, 24, 22, 9, 8, 2, 14, 15, 1, 11, 12, 3, 23, 25, 21, 5, 19, 7, 17, 6, 18},
	}

	UKW_c := Rotor{
		Name:  "Narrow UKW-c",
		Wires: []byte{17, 3, 14, 1, 9, 13, 19, 10, 21, 4, 7, 12, 11, 5, 2, 22, 25, 0, 23, 6, 24, 8, 15, 18, 20, 16},
	}

	set = RotorSet{
		I:     I,
		II:    II,
		III:   III,
		IV:    IV,
		V:     V,
		VI:    VI,
		VII:   VII,
		VIII:  VIII,
		UKW_A: UKW_A,
		UKW_B: UKW_B,
		UKW_C: UKW_C,
		UKW_b: UKW_b,
		UKW_c: UKW_c,
		Beta:  Beta,
		Gamma: Gamma,
	}

	return set
}
