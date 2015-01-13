package chance

const (
	chars      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	numChars   = "0123456789"
	upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

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

