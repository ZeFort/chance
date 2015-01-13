package chance

func (chance *Chance) String() string {
	var buffer bytes.Buffer
	l := chance.Natural()
	for i := 0; i < l; i++ {
		buffer.WriteString(chance.Char())
	}
	return buffer.String()
}

func (chance *Chance) AnyString(str string) string {
	var buffer bytes.Buffer
	l := chance.Natural()
	for i := 0; i < l; i++ {
		buffer.WriteString(chance.AnyChar(str))
	}
	return buffer.String()
}
