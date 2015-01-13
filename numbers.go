package chance

import "math"

func (chance *Chance) Float() float64 {
	sign := 1
	if (chance.Bool()) {
		sign = -1
	}
	return math.MaxInt64 * chance.r.Float64() * sign
}

func (chance *Chance) FloatBtw(min float64, max float64) float64 {
	return chance.r.Float64() * (max - min) + min
}

func (chance *Chance) IntN(max int) int {
	return chance.r.Intn(max)
}

func (chance *Chance) Natural() int {
	return chance.r.Intn() + 1
}
