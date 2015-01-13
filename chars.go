package chance

const (
	chars      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	numChars   = "0123456789"
	upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerChars = "abcdefghijklmnopqrstuvwxyz"
)

func (chance *Chance) AlphaChar() uint8 {
	l := len(alphaChars)
	return alphaChars[chance.NaturalN(l)]
}

func (chance *Chance) LowerChar() uint8 {
	l := len(lowerChars)
	return lowerChars[chance.NaturalN(l)]
}

func (chance * Chance) Char() uint8 {
	l := len(chars)
	return chars[chance.NaturalN(l)]
}

func (chance * Chance) NumChar() uint8 {
	l := len(numChars)
	return numChars[chance.NaturalN(l)]
}

func (chance * Chance) UpperChar() uint8 {
	l := len(upperChars)
	return upperChars[chance.NaturalN(l)]
}

func (chance * Chance) AnyChar(str string) uint8{
	l := len(str)
	return str[chance.NaturalN(l)]
}

