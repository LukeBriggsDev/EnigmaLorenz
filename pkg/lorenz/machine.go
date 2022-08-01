package lorenz

import "EnigmaLorenz/pkg/util"

func boolToByte(b bool) byte {
	if b {
		return 1
	} else {
		return 0
	}
}

type Wheel struct {
	pins []bool
	pos  byte
}

func NewWheel(pins []bool, pos byte) Wheel {
	return Wheel{
		pins: pins,
		pos:  pos,
	}
}

func (w *Wheel) rotate() {
	w.pos = byte(util.NegMod(int(w.pos)-1, len(w.pins)))
}

func (w *Wheel) getCurrentPin() bool {
	return w.pins[w.pos]
}

func (w *Wheel) GetPins() []bool {
	return w.pins
}

type Lorenz struct {
	chiWheels []Wheel
	m1        Wheel
	m2        Wheel
	psiWheels []Wheel
}

func NewLorenz(chiWheels []Wheel, m1 Wheel, m2 Wheel, psiWheels []Wheel) Lorenz {
	return Lorenz{
		chiWheels: chiWheels,
		m1:        m1,
		m2:        m2,
		psiWheels: psiWheels,
	}
}

func (m *Lorenz) Encrypt(plain []byte) []byte {
	ciphertext := []byte{}

	for _, char := range plain {
		key := byte(0)
		// Apply  wheels
		key = WheelsToByte(m.chiWheels)
		psi := WheelsToByte(m.psiWheels)
		key = key ^ psi
		ciphertext = append(ciphertext, char^key)
		// Rotate chi wheels
		for i := 0; i < len(m.chiWheels); i++ {
			m.chiWheels[i].rotate()
		}

		// Rotate motor wheels
		m.m1.rotate()
		if m.m1.getCurrentPin() {
			m.m2.rotate()
		}

		// Rotate psi wheels
		if m.m2.getCurrentPin() {
			for i := 0; i < len(m.psiWheels); i++ {
				m.psiWheels[i].rotate()
			}
		}

	}
	return ciphertext
}

func WheelsToByte(w []Wheel) byte {
	result := byte(0)
	for i := 0; i < len(w); i++ {
		result += boolToByte(w[i].getCurrentPin()) << byte(len(w)-i-1)
	}
	return result
}

type WheelSet struct {
	Chi1 Wheel
	Chi2 Wheel
	Chi3 Wheel
	Chi4 Wheel
	Chi5 Wheel
	M1   Wheel
	M2   Wheel
	Psi1 Wheel
	Psi2 Wheel
	Psi3 Wheel
	Psi4 Wheel
	Psi5 Wheel
}

func NewWheelSet() WheelSet {
	return WheelSet{
		Wheel{
			pins: []bool{false, false, false, true, true, true, true, false, false, false, false, true, true, false, false, false, false, true, false, true, true, false, false, true, false, false, true, true, false, true, false, true, true, false, true, false, true, true, true, true, false},
			pos:  0,
		},
		Wheel{
			pins: []bool{true, true, false, false, true, true, true, false, true, true, false, false, false, true, false, true, false, true, true, false, false, false, true, false, false, false, false, true, true, true, false},
			pos:  0,
		},
		Wheel{
			pins: []bool{false, false, true, true, true, false, true, true, false, false, true, false, false, false, false, true, true, true, false, false, true, true, false, true, true, false, false, true, true},
			pos:  0,
		},
		Wheel{
			pins: []bool{false, false, true, true, false, false, true, false, true, true, false, false, true, false, false, true, true, false, false, true, false, false, true, true, true, true},
			pos:  0,
		},
		Wheel{
			pins: []bool{false, true, false, false, false, true, false, true, true, false, false, true, false, false, false, true, true, true, false, true, true, true, false},
			pos:  0,
		},
		Wheel{
			pins: []bool{true, true, true, false, true, false, true, true, false, false, true, true, false, false, true, true, false, false, false, true, true, true, true, false, true, false, true, true, false, true, true, false, false, false, true, true, false, false, false, false, true, true, true, true, false, true, true, false, false, true, true, false, false, false, true, true, false, false, false, false, true},
			pos:  0,
		},
		Wheel{
			pins: []bool{true, false, true, true, true, false, true, false, true, false, true, false, true, false, false, true, false, true, false, true, true, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false},
			pos:  0,
		},
		Wheel{
			pins: []bool{false, false, true, false, true, false, true, false, true, false, true, false, false, true, false, false, true, false, true, true, false, true, true, false, true, false, true, false, false, true, true, false, true, true, true, false, false, true, true, true, false, false, false},
			pos:  0,
		},
		Wheel{
			pins: []bool{false, false, true, false, true, true, false, true, false, true, false, true, false, true, false, true, false, true, true, false, false, true, true, false, true, false, false, true, false, true, true, true, true, false, false, false, false, false, true, true, true, false, false, true, false, true, true},
			pos:  0,
		},
		Wheel{
			pins: []bool{true, false, true, false, true, false, true, false, true, false, true, false, true, false, false, true, false, false, true, true, false, true, false, true, false, true, true, true, true, false, false, false, false, true, true, true, false, false, false, true, true, true, false, true, true, false, false, true, false, false, true},
			pos:  0,
		},
		Wheel{
			pins: []bool{true, false, true, false, false, true, true, false, true, false, true, false, true, false, true, false, true, false, true, true, false, true, false, false, false, false, true, true, false, false, true, true, false, false, true, true, false, true, true, true, true, true, false, true, false, false, true, false, false, false, false, true, false},
			pos:  0,
		},
		Wheel{
			pins: []bool{false, true, false, true, false, true, false, true, false, true, true, false, false, false, true, false, true, false, false, true, true, true, false, true, true, true, true, false, true, true, false, true, false, false, false, false, true, false, false, false, true, false, false, true, true, false, true, true, false, false, true, true, false, false, true, false, true, false, true},
			pos:  0,
		},
	}
}
