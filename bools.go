package chance

func (chance *Chance) BoolWithChance(likeliness int) bool {
	if (likeliness < 0) || (likeliness > 100) {
		return false
	}
	f := chance.r.Float64() * 100
	return likeliness <= f
}
