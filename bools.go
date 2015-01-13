package chance

func (chance *Chance) BoolWithChance(likeliness int) bool {
	if (likeliness < 0) || (likeliness > 100) {
		return false
	}
	f := chance.r.Float64() * 100
	return float64(likeliness) <= f
}

func (chance *Chance) Bool() bool {
	return chance.r.Intn(2) == 1
}
