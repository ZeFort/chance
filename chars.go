package chance

const (
	chars      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	numChars   = "0123456789"
	upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerChars = "abcdefghijklmnopqrstuvwxyz"
	vowelChars = "aeoui"
	consonantChars = "bcdfghjklmnpqrstvwxyz"
)

func (chance *Chance) AlphaChar() string {
	l := len(alphaChars)
	return string(alphaChars[chance.NaturalN(l)])
}

func (chance *Chance) LowerChar() string {
	l := len(lowerChars)
	return string(lowerChars[chance.NaturalN(l)])
}

func (chance *Chance) Char() string {
	l := len(chars)
	return string(chars[chance.NaturalN(l)])
}

func (chance *Chance) NumChar() string {
	l := len(numChars)
	return string(numChars[chance.NaturalN(l)])
}

func (chance *Chance) UpperChar() string {
	l := len(upperChars)
	return string(upperChars[chance.NaturalN(l)])
}

func (chance *Chance) AnyChar(str string) string {
	l := len(str)
	return string(str[chance.NaturalN(l)])
}

func (chance *Chance) VowelChar() string {
	return string(vowelChars[chance.NaturalN(len(vowelChars))])
}

func (chance *Chance) ConsonantChar() string {
	return string(consonantChars[chance.NaturalN(len(consonantChars))])
}
