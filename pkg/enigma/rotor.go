// Package enigma implements the necessary functions to simulate all Enigma machines from Enigma 1 to Enigma M4.
package enigma

import (
	"errors"
	"log"
)

// indexOf is a helper function to find the location of a byte in a byte slice
func indexOf(needle byte, haystack []byte) (byte, error) {
	for idx, val := range haystack {
		if val == needle {
			return byte(idx), nil
		}
	}
	return 0, errors.New("Item not found in list")
}

// negMod is a helper function to have a mod n produce a positive result even if a is negative.
func negMod(a int, n int) int {
	return (a%n + n) % n
}

// A Rotor contains all the information necessary to translate an input connection to an output connection.
// The Name of the rotor contains the string representation.
// In the standard RotorSet this is in the form of either a Roman numeral or a Greek letter.
//
// Wires is a list of mappings of the input to the output.
// The index of the list can be seen as the input and the value can be seen as the output.
// e.g.
//    | Index | 0(A)  1(B)  2(C) 3(D)  4(E) 5(F) 6(G)  7(H)  8(I)  9(J) 10(K) 11(L) 12(M) 13(N) 14(O) 15(P) 16(Q) 17(R) 18(S) 19(T) 20(U) 21(V) 22(W) 23(X) 24(Y) 25(Z)
//    | Value | 4(E) 10(K) 12(M) 5(F) 11(L) 6(G) 3(D) 16(Q) 21(V) 25(Z) 13(N) 19(T) 14(O) 22(W) 24(Y)  7(H) 23(X) 20(U) 18(S) 15(P)  0(A)  8(I)  1(B) 17(R)  2(C)  9(J)
//
// The wiring for the standard RotorSet was acquired from
// https://www.cryptomuseum.com/crypto/enigma/m4/index.htm#Wiring
//
// TurnoverList contains the list of values that will be the current value when notch is aligned with the mechanism.
// When a value in the turnover list is the current position,
// the Enigma Encrypt function will step the rotor to its left on the next rotation.
type Rotor struct {
	Name         string
	Wires        []byte
	shownPos     byte
	TurnoverList []byte
	ringSetting  byte
}

// SetShownPos is analogous to setting the rotor position by changing the letter/number shown in the window.
// The position can be any number between 1 and 26 inclusively.
//
// Errors
//
// A fatal error will occur if a value less than 0 or more than 26 is passed in.
func (r *Rotor) SetShownPos(pos byte) {
	if pos < 1 || pos > 26 {
		log.Fatal("Rotor position must be set to value between 1 and 26")
	}
	r.shownPos = pos - 1
}

// GetShownPos will return the value that would be showing through the Enigma window.
// The value returned will be a value between 1 and 26 inclusively.
func (r *Rotor) GetShownPos() byte {
	return r.shownPos + 1
}

// SetRingSetting is equivalent to the ring setting on a real Enigma.
// It provides an offset between the input wiring and output wiring, achieved by rotating one side of the rotor wiring.
// The offset can be any number between 0 and 25 inclusively.
//
// Errors
//
// A fatal error will occur if a value less then 0 or greater than 25 is passed as a parameter.
func (r *Rotor) SetRingSetting(offset byte) {
	if offset < 0 || offset > 25 {
		log.Fatal("Rotor position must be set to value between 0 and 25")
	}
	r.ringSetting = offset
}

// GetRingSetting will return the current ring setting.
func (r *Rotor) GetRingSetting() byte {
	return r.ringSetting
}

func (r *Rotor) normalizedPos() byte {
	return byte(negMod(int(r.shownPos)-int(r.ringSetting), len(r.Wires)))
}

// AtNotch will return whether the rotor is currently in the position where its notch aligns with the pawl.
func (r *Rotor) AtNotch() bool {
	_, err := indexOf(r.shownPos, r.TurnoverList)
	return err == nil
}

// Rotate will rotate the rotor counterclockwise by one step.
func (r *Rotor) Rotate() {
	r.shownPos += 1
	r.shownPos %= byte(len(r.Wires))
}

// Translate will return the output signal from a given input signal.
func (r Rotor) Translate(plain byte) byte {

	shiftedPlain := (plain + r.normalizedPos()) % byte(len(r.Wires)) // Shift by rotor and ring position.
	shiftedCipher := negMod(int(r.Wires[shiftedPlain])-int(r.normalizedPos()), len(r.Wires))
	return byte(shiftedCipher)
}

// TranslateReverse will return the input signal for a given output signal.
func (r Rotor) TranslateReverse(cipher byte) byte {
	unshiftedCipher := (cipher + r.normalizedPos()) % byte(len(r.Wires))
	unshiftedPlain, _ := indexOf(unshiftedCipher, r.Wires)
	shiftedPlain := negMod(int(unshiftedPlain)-int(r.normalizedPos()), len(r.Wires))
	return byte(shiftedPlain)
}

// RotorSet contains all the standard rotors and reflectors that were available from Enigma 1 to M4 Enigma.
// The rotors were gathered from [Crypto Museum]: https://www.cryptomuseum.com/crypto/enigma/m4/index.htm#wiring
//
// Further Notes
//
// 4 rotor enigma only allowed for a specific set of rotors to be used as the 4th rotor.
// These rotors are Beta and Gamma.
// UKW_b and UKW_c were designed so that when UKW_b and Beta, or UKW_c and Gamma were used together in default positions,
// they were equivalent to UKW_B and UKW_C respectively
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

// GenerateRotors returns the standard RotorSet
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
