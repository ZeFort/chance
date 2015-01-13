package chance

import (
	"bytes"
)

func (chance *Chance) String() string {
	var buffer bytes.Buffer
	l := chance.Natural()
	for i := 0; i < l; i++ {
		buffer.WriteString(string(chance.Char()))
	}
	return buffer.String()
}

func (chance *Chance) StringN(l int) string {
	var buffer bytes.Buffer
	for i := 0; i < l; i++ {
		buffer.WriteString(string(chance.Char()))
	}
	return buffer.String()
}

func (chance *Chance) AnyString(str string) string {
	var buffer bytes.Buffer
	l := chance.Natural()
	for i := 0; i < l; i++ {
		buffer.WriteString(string(chance.AnyChar(str)))
	}
	return buffer.String()
}

func (chance *Chance) AnyStringN(str string, l int) string {
	var buffer bytes.Buffer
	for i := 0; i < l; i++ {
		buffer.WriteString(string(chance.AnyChar(str)))
	}
	return buffer.String()
}
