package lorenz

import "errors"

type bimap struct {
	forwardMap map[byte]byte
	reverseMap map[byte]byte
}

func bimapFromSlice(s []byte) bimap {
	m := bimap{
		forwardMap: make(map[byte]byte),
		reverseMap: make(map[byte]byte),
	}
	for idx, val := range s {
		m.Add(byte(idx), val)
	}
	return m
}

func (m *bimap) Add(k byte, v byte) {
	m.forwardMap[k] = v
	m.reverseMap[v] = k
}

// GetITA2Code returns the corresponding ITA2 character for an ASCII character.
func (m *bimap) GetITA2Code(k byte) (byte, bool) {
	v, Ok := m.reverseMap[k]
	return v, Ok
}

// GetASCII returns the corresponding ASCII character for an ITA2 character.
func (m *bimap) GetASCII(k byte) (byte, bool) {
	v, Ok := m.forwardMap[k]
	return v, Ok
}

type ITA2 struct {
	letterAlphabet bimap
	figureAlphabet bimap
	figShift       byte
	letShift       byte
}

// NewITA2LSB creates an ITA2 struct with the corresponding ITA2 alphabets with the least significant bit on the left.
func NewITA2LSB() ITA2 {
	letter := []byte{
		'±', // NULL
		'T',
		'_', // CR
		'O',
		' ',
		'H',
		'N',
		'M',
		'|', // LF
		'L',
		'R',
		'G',
		'I',
		'P',
		'C',
		'V',
		'E',
		'Z',
		'D',
		'B',
		'S',
		'Y',
		'F',
		'X',
		'A',
		'W',
		'J',
		'^', // Shift In
		'U',
		'Q',
		'K',
		'*', // Shift Out
	}

	figure := []byte{
		'±', // NULL
		'5',
		'_', // LF
		'9',
		'!',
		'£',
		',',
		'.',
		'|', // CR
		')',
		'4',
		'&',
		'8',
		'0',
		':',
		'=',
		'3',
		'+',
		'#', // Enquiry
		'?',
		'\'',
		'6',
		'%',
		'/',
		'-',
		'2',
		'$', // Bell
		'^', // Shift In
		'7',
		'1',
		'(',
		'*', // Shift Out
	}

	ITA := ITA2{
		letterAlphabet: bimapFromSlice(letter),
		figureAlphabet: bimapFromSlice(figure),
		figShift:       '^',
		letShift:       '*',
	}

	return ITA
}

// AsciiToITA2 takes a string s and returns a slice of the translated ITA2 bytes along with an error
//
// # Errors
//
// An error will be returned if one of the characters in the string does not appear in the ITA2 alphabet.
func (alphabet *ITA2) AsciiToITA2(s string, decrypt bool) ([]byte, error) {
	encoded := []byte{}
	inLetterShift := true
	for _, char := range s {
		letter, letterExist := alphabet.letterAlphabet.GetITA2Code(byte(char))
		if !letterExist {
			figure, figureExists := alphabet.figureAlphabet.GetITA2Code(byte(char))
			if !figureExists {
				return []byte(""), errors.New("invalid characters in input string")
			}
			if inLetterShift && !decrypt {
				itaFig, _ := alphabet.letterAlphabet.GetITA2Code(alphabet.figShift)
				encoded = append(encoded, itaFig)
				inLetterShift = !inLetterShift
			}
			encoded = append(encoded, figure)
		} else {
			if !inLetterShift && !decrypt {
				itaLet, _ := alphabet.letterAlphabet.GetITA2Code(alphabet.letShift)
				encoded = append(encoded, itaLet)
				inLetterShift = !inLetterShift
			}
			encoded = append(encoded, letter)
		}
	}
	return encoded, nil
}

// ITA2ToAscii takes a slice of ITA2 bytes and returns a string of ASCII characters along with an error.
//
// # Errors
//
// An error will be returned if there is no corresponding ASCII character for the ITA2 byte
func (alphabet *ITA2) ITA2ToAscii(b []byte, decrypt bool) (string, error) {
	decoded := ""
	inLetterShift := true
	for _, char := range b {
		figShiftByte, _ := alphabet.letterAlphabet.GetITA2Code(alphabet.figShift)
		letShiftByte, _ := alphabet.letterAlphabet.GetITA2Code(alphabet.letShift)
		if char == figShiftByte {
			inLetterShift = false
			if decrypt {
				continue
			}

		} else if char == letShiftByte {
			inLetterShift = true
			if decrypt {
				continue
			}
		}

		var plain byte
		var plainExist bool

		if inLetterShift {
			plain, plainExist = alphabet.letterAlphabet.GetASCII(char)
		} else {
			plain, plainExist = alphabet.figureAlphabet.GetASCII(char)
		}

		if !plainExist {
			return "", errors.New("incorrect character sequence")
		}
		decoded += string(plain)

	}

	return decoded, nil

}
