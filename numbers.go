package chance

func (chance *Chance) RInt() int {
	return chance.r.Intn(1)
}
