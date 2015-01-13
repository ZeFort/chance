package chance

const (
	chars      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	numChars   = "0123456789"
	upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerChars = "abcdefghijklmnopqrstuvwxyz"
)

func (chance *Chance) AlphaChar() int {
	l := len(alphaChars)
	return alphaChars[chance.r.NaturalN(l)]
}

func (chance *Chance) LowerChar() int {
	l := len(lowerChars)
	return lowerChars[chance.r.NaturalN(l)]
}

func (chance * Chance) Char() {
	l := len(chars)
	return chars[chance.r.NaturalN(l)]
}

func (chance * Chance) NumChar() {
	l := len(numChars)
	return numChars[chance.r.NaturalN(l)]
}

func (chance * Chance) UpperChar() {
	l := len(upperChars)
	return upperChars[chance.r.NaturalN(l)]
}

func (chance * Chance) AnyChar(str string) {
	l := len(str)
	return str[chance.r.NaturalN(l)]
}

