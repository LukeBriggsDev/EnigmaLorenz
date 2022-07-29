package lorenz

import "fmt"

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
	w.pos = (w.pos + 1) % byte(len(w.pins))
}

func (w *Wheel) getCurrentPin() bool {
	return w.pins[w.pos]
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

func (m *Lorenz) Encrypt(plain []byte) {
	//ciphertext := []byte{}

	for _, char := range plain {
		key := byte(0)
		// Apply  wheels
		key = WheelsToByte(m.chiWheels)
		key = key ^ WheelsToByte(m.psiWheels)

		// Rotate chi wheels
		for _, wheel := range m.chiWheels {
			wheel.rotate()
		}

		// Rotate motor wheels
		m.m1.rotate()
		if m.m1.getCurrentPin() {
			m.m2.rotate()
		}

		// Rotate psi wheels
		if m.m2.getCurrentPin() {
			for _, wheel := range m.psiWheels {
				wheel.rotate()
			}
		}

		fmt.Println(key, char)
	}
}

func WheelsToByte(w []Wheel) byte {
	result := byte(0)
	for i := 0; i < len(w); i++ {
		result += boolToByte(w[i].getCurrentPin()) << byte(len(w)-i-1)
	}
	return result
}
