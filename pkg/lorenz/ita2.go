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
}

// NewITA2LSB creates an ITA2 struct with the corresponding ITA2 alphabets with the least significant bit on the left.
func NewITA2LSB() ITA2 {
	letter := []byte{
		'\x00', // NULL
		'T',
		'\x0A', // CR
		'O',
		' ',
		'H',
		'N',
		'M',
		'\x0D', // LF
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
		'\x1B', // Shift In
		'U',
		'Q',
		'K',
		'\x1F', // Shift Out
	}

	figure := []byte{
		'\x00', // NULL
		'5',
		'\x0A', // LF
		'9',
		' ',
		'£',
		',',
		'.',
		'\x0D', // CR
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
		'\x07', // Bell
		'\x1B', // Shift In
		'7',
		'1',
		'(',
		'\x1F', // Shift Out
	}

	ITA := ITA2{
		letterAlphabet: bimapFromSlice(letter),
		figureAlphabet: bimapFromSlice(figure),
	}

	return ITA
}

// AsciiToITA2 takes a string s and returns a slice of the translated ITA2 bytes along with an error
//
// Errors
//
// An error will be returned if one of the characters in the string does not appear in the ITA2 alphabet.
func (alphabet *ITA2) AsciiToITA2(s string) ([]byte, error) {
	encoded := []byte{}
	inLetterShift := true
	for _, char := range s {
		letter, letterExist := alphabet.letterAlphabet.GetITA2Code(byte(char))
		if !letterExist {
			figure, figureExists := alphabet.figureAlphabet.GetITA2Code(byte(char))
			if !figureExists {
				return []byte(""), errors.New("invalid characters in input string")
			}
			if inLetterShift {
				encoded = append(encoded, 0x1B)
				inLetterShift = !inLetterShift
			}
			encoded = append(encoded, figure)
		} else {
			if !inLetterShift {
				encoded = append(encoded, 0x1F)
				inLetterShift = !inLetterShift
			}
			encoded = append(encoded, letter)
		}
	}
	return encoded, nil
}

// ITA2ToAscii takes a slice of ITA2 bytes and returns a string of ASCII characters along with an error.
//
// Errors
//
// An error will be returned if there is no corresponding ASCII character for the ITA2 byte
//
func (alphabet *ITA2) ITA2ToAscii(b []byte) (string, error) {
	decoded := ""
	inLetterShift := true
	for _, char := range b {
		if char == 0x1B {
			inLetterShift = false
			continue
		} else if char == 0x1F {
			inLetterShift = true
			continue
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
