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
	chiWheels   [5]Wheel
	motorWheels [2]Wheel
	psiWheels   [5]Wheel
}

func NewLorenz(chiWheels [5]Wheel, motorWheels [2]Wheel, psiWheels [5]Wheel) Lorenz {
	return Lorenz{
		chiWheels:   chiWheels,
		motorWheels: motorWheels,
		psiWheels:   psiWheels,
	}
}

// ResetRotorPos sets the rotor positions of all the machine's rotors to 0.
func (m *Lorenz) ResetRotorPos() {
	for i := 0; i < len(m.chiWheels); i++ {
		m.chiWheels[i].pos = 0
	}
	for i := 0; i < len(m.psiWheels); i++ {
		m.psiWheels[i].pos = 0
	}
	for i := 0; i < len(m.motorWheels); i++ {
		m.motorWheels[i].pos = 0
	}
}

// SetChiPos takes an array of 5 rotor positions and sets the corresponding Chi rotors to them.
func (m *Lorenz) SetChiPos(positions [5]byte) {
	for i := 0; i < len(m.chiWheels); i++ {
		m.chiWheels[i].pos = positions[i]
	}
}

// SetPsiPos takes an array of 5 rotor positions and sets the corresponding Psi rotors to them.
func (m *Lorenz) SetPsiPos(positions [5]byte) {
	for i := 0; i < len(m.psiWheels); i++ {
		m.psiWheels[i].pos = positions[i]
	}
}

// SetMotorPos takes an array of 2 rotor positions and sets the corresponding Motor rotors to them.
func (m *Lorenz) SetMotorPos(positions [2]byte) {
	for i := 0; i < len(m.motorWheels); i++ {
		m.motorWheels[i].pos = positions[i]
	}
}

// Encrypt takes a slice of bytes and returns the result of them passing through the Lorenz machine.
func (m *Lorenz) Encrypt(plain []byte) []byte {
	ciphertext := []byte{}

	for _, char := range plain {
		key := byte(0)
		// Apply  wheels
		key = WheelsToByte(m.chiWheels[:])
		psi := WheelsToByte(m.psiWheels[:])
		key = key ^ psi
		ciphertext = append(ciphertext, char^key)
		// Rotate chi wheels
		for i := 0; i < len(m.chiWheels); i++ {
			m.chiWheels[i].rotate()
		}

		// Rotate psi wheels
		if m.motorWheels[1].getCurrentPin() {
			for i := 0; i < len(m.psiWheels); i++ {
				m.psiWheels[i].rotate()
			}
		}

		// Rotate motor wheels
		m.motorWheels[0].rotate()
		if m.motorWheels[0].getCurrentPin() {
			m.motorWheels[1].rotate()
		}

	}
	return ciphertext
}

// WheelsToByte takes a list of wheels and uses the current pin to create a byte with each bit mapping to each pin.
func WheelsToByte(w []Wheel) byte {
	result := byte(0)
	for i := 0; i < len(w); i++ {
		result += boolToByte(w[i].getCurrentPin()) << byte(len(w)-i-1)
	}
	return result
}

type WheelSet struct {
	Chi   [5]Wheel
	Motor [2]Wheel
	Psi   [5]Wheel
}

func NewWheelSet() WheelSet {
	return WheelSet{
		[5]Wheel{
			{
				pins: []bool{false, true, true, false, true, true, false, false, false, true, true, false, true, true, false, false, true, false, false, false, false, true, true, true, false, false, true, true, true, false, false, false, false, true, true, true, false, false, true, true, false},
				pos:  0,
			},
			{
				pins: []bool{true, true, false, true, true, false, false, false, false, true, true, true, false, true, true, true, true, false, true, false, false, false, true, true, false, false, true, true, false, false, false},
				pos:  0,
			},
			{
				pins: []bool{false, false, true, false, false, true, true, false, false, false, true, true, false, false, false, true, true, true, false, false, false, true, true, false, true, true, true, true, false},
				pos:  0,
			},
			{
				pins: []bool{true, false, true, false, true, false, false, true, true, false, false, false, true, true, false, false, true, false, true, true, true, false, false, true, false, true},
				pos:  0,
			},
			{
				pins: []bool{false, true, false, false, true, true, true, true, false, false, false, true, false, true, true, true, false, false, false, false, true, false, true},
				pos:  0,
			},
		},
		[2]Wheel{
			{
				pins: []bool{true, false, true, true, false, true, false, true, true, true, false, true, true, true, false, true, false, true, false, true, true, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true, true, false, true, true, true, false, true, true, true, false, true, false, true, false, true, true, true, true, false, true, false, true, false, true},
				pos:  0,
			},
			{
				pins: []bool{false, true, false, true, false, true, true, false, true, false, true, true, false, true, true, true, false, true, true, true, false, true, true, false, true, false, true, true, true, false, true, true, true, false, true, true, true},
				pos:  0,
			},
		},
		[5]Wheel{
			{
				pins: []bool{true, true, false, true, false, false, true, true, false, false, false, true, true, true, false, false, true, true, false, false, false, true, true, false, false, false, true, true, true, true, false, false, true, true, true, false, false, true, true, true, false, false, false},
				pos:  0,
			},
			{
				pins: []bool{false, false, false, true, false, false, false, true, true, true, false, false, true, true, false, false, true, true, true, false, false, false, true, true, true, true, false, false, false, true, true, false, false, true, true, true, false, false, true, true, true, false, false, true, false, true, true},
				pos:  0,
			},
			{
				pins: []bool{false, true, false, false, true, true, false, false, true, true, true, false, false, true, true, true, false, false, true, false, false, false, true, true, true, true, false, false, false, true, false, false, false, true, true, true, false, false, false, true, true, false, false, false, true, true, false, false, true, true, true},
				pos:  0,
			},
			{
				pins: []bool{false, false, true, true, true, false, false, true, true, false, false, true, true, true, false, false, true, true, true, true, false, false, false, true, false, false, false, true, true, false, false, true, true, true, false, false, true, false, false, true, true, false, false, false, true, true, false, false, true, true, true, false, true},
				pos:  0,
			},
			{
				pins: []bool{true, false, false, true, true, true, false, false, false, true, false, false, false, true, true, true, true, false, false, true, true, true, false, false, true, false, false, true, true, true, true, false, false, false, true, true, false, false, true, true, true, false, false, true, true, false, false, true, true, true, false, false, true, false, false, false, true, true, false},
				pos:  0,
			},
		},
	}
}
