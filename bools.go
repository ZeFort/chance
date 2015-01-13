package chance

func (chance *Chance) BoolWithChance(likeliness int) bool {
	if (likeliness < 0) || (likeliness > 100) {
		return false
	}
	f := chance.r.Float64() * 100
	return likeliness <= f
}

func (chance *Chance) Bool() int {
	if (chance.r.Intn(1) == 1){
		return true
	}else{
		return false
	}
}
