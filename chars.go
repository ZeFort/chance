package chance

const (
	chars          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	numChars       = "0123456789"
	upperChars     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaChars     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerChars     = "abcdefghijklmnopqrstuvwxyz"
	vowelChars     = "aeoui"
	consonantChars = "bcdfghjklmnpqrstvwxyz"
)

// AlphaChar returns any char from the next list: `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
func (chance *Chance) AlphaChar() string {
	l := len(alphaChars)
	return string(alphaChars[chance.IntN(l)])
}

// LowelChar returns any lowercase char from the next list: `abcdefghijklmnopqrstuvwxyz`
func (chance *Chance) LowerChar() string {
	l := len(lowerChars)
	return string(lowerChars[chance.IntN(l)])
}

// Char returns any char from the next list: `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`
func (chance *Chance) Char() string {
	l := len(chars)
	return string(chars[chance.IntN(l)])
}

// NumChar returns random numeric char
func (chance *Chance) NumChar() string {
	l := len(numChars)
	return string(numChars[chance.IntN(l)])
}

// UpperChar returns any uppercase char from the next list: `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
func (chance *Chance) UpperChar() string {
	l := len(upperChars)
	return string(upperChars[chance.IntN(l)])
}

// AnyChar returns random char from string passed as argument
func (chance *Chance) AnyChar(str string) string {
	l := len(str)
	return string(str[chance.IntN(l)])
}

// VowelChar returns random vowel char from the next list: `aeoui`
func (chance *Chance) VowelChar() string {
	l := len(vowelChars)
	return string(vowelChars[chance.IntN(l)])
}

// ConsonantChar returns random consonant char from the next list: `bcdfghjklmnpqrstvwxyz`
func (chance *Chance) ConsonantChar() string {
	l := len(consonantChars)
	return string(consonantChars[chance.IntN(l)])
}
