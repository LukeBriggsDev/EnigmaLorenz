package enigma

type Plugboard struct {
	state map[byte]byte
}

func NewPlugboard() Plugboard {
	return Plugboard{
		state: make(map[byte]byte),
	}
}

func (p *Plugboard) AddPlug(letter1 byte, letter2 byte) {
	p.state[letter1] = letter2
	p.state[letter2] = letter1
}

func (p *Plugboard) RemovePlug(letter byte) {
	val, exists := p.state[letter]
	if !exists {
		return
	}
	delete(p.state, val)
	delete(p.state, letter)
}

func (p *Plugboard) Translate(letter byte) byte {
	val, exists := p.state[letter]
	if !exists {
		return letter
	}
	return val
}
