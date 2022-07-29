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

func (m *bimap) GetITA2Code(k byte) (byte, bool) {
	v, Ok := m.reverseMap[k]
	return v, Ok
}

func (m *bimap) GetASCII(k byte) (byte, bool) {
	v, Ok := m.forwardMap[k]
	return v, Ok
}

type ITA2 struct {
	letterAlphabet bimap
	figureAlphabet bimap
}

func NewITA2() ITA2 {
	letter := []byte{
		'\x00', // NULL
		'E',
		'\x0A', // LF
		'A',
		' ',
		'S',
		'I',
		'U',
		'\x0D', // CR
		'D',
		'R',
		'J',
		'N',
		'F',
		'C',
		'K',
		'T',
		'Z',
		'L',
		'W',
		'H',
		'Y',
		'P',
		'Q',
		'O',
		'B',
		'G',
		'\x1B', // Shift In
		'M',
		'X',
		'V',
		'\x1F', // Shift Out
	}

	figure := []byte{
		'\x00', // NULL
		'3',
		'\x0A', // LF
		'-',
		' ',
		'\'',
		'8',
		'7',
		'\x0D', // CR
		'\x05', // Enquiry
		'4',
		'\x07', // Bell
		',',
		'!',
		':',
		'(',
		'5',
		'+',
		')',
		'2',
		'$',
		'6',
		'0',
		'1',
		'9',
		'?',
		'&',
		'\x1B', // Shift In
		'.',
		'/',
		';',
		'\x1F', // Shift Out
	}

	ITA := ITA2{
		letterAlphabet: bimapFromSlice(letter),
		figureAlphabet: bimapFromSlice(figure),
	}

	return ITA
}

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

		if inLetterShift {
			letter, letterExists := alphabet.letterAlphabet.GetASCII(char)
			if !letterExists {
				return "", errors.New("incorrect character sequence")
			}
			decoded += string(letter)
		} else {
			figure, figureExists := alphabet.figureAlphabet.GetASCII(char)
			if !figureExists {
				return "", errors.New("incorrect character sequence")
			}
			decoded += string(figure)
		}

	}

	return decoded, nil

}
