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
	return string(alphaChars[chance.IntN(l)])
}

func (chance *Chance) LowerChar() string {
	l := len(lowerChars)
	return string(lowerChars[chance.IntN(l)])
}

func (chance *Chance) Char() string {
	l := len(chars)
	return string(chars[chance.IntN(l)])
}

func (chance *Chance) NumChar() string {
	l := len(numChars)
	return string(numChars[chance.IntN(l)])
}

func (chance *Chance) UpperChar() string {
	l := len(upperChars)
	return string(upperChars[chance.IntN(l)])
}

func (chance *Chance) AnyChar(str string) string {
	l := len(str)
	return string(str[chance.IntN(l)])
}

func (chance *Chance) VowelChar() string {
	l := len(vowelChars)
	return string(vowelChars[chance.IntN(l)])
}

func (chance *Chance) ConsonantChar() string {
	l:= len(consonantChars)
	return string(consonantChars[chance.IntN(l)])
}
