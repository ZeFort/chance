package chance

// BoolWithChance returns any boolean value, where `true` value is returned with some likeliness.
func (chance *Chance) BoolWithChance(likeliness int) bool {
	if (likeliness < 0) || (likeliness > 100) {
		return false
	}
	f := chance.r.Float64() * 100
	return float64(likeliness) <= f
}

// Bool returns any boolean value, where `true` and `false` has equivalent chance.
func (chance *Chance) Bool() bool {
	return chance.BoolWithChance(50)
}
