package chance

import (
	"bytes"
)

// String returns random string with length in range [1..100]
func (chance *Chance) String() string {
	var buffer bytes.Buffer
	l := chance.NaturalN(100)
	for i := 0; i < l; i++ {
		buffer.WriteString(string(chance.Char()))
	}
	return buffer.String()
}

// String returns random string with length in range [1..len]
func (chance *Chance) StringN(len int) string {
	var buffer bytes.Buffer
	for i := 0; i < len; i++ {
		buffer.WriteString(string(chance.Char()))
	}
	return buffer.String()
}

// AnyString returns random string with length in range [1..100] based on symbols of string passed as argument.
func (chance *Chance) AnyString(str string) string {
	var buffer bytes.Buffer
	l := chance.NaturalN(100)
	for i := 0; i < l; i++ {
		buffer.WriteString(string(chance.AnyChar(str)))
	}
	return buffer.String()
}

// String returns random string with length in range [1..len] based on symbols of string passed as argument.
func (chance *Chance) AnyStringN(str string, len int) string {
	var buffer bytes.Buffer
	for i := 0; i < len; i++ {
		buffer.WriteString(string(chance.AnyChar(str)))
	}
	return buffer.String()
}
