package chance

func (chance *Chance) RInt() int64 {
	return chance.r.Intn(10)
}
