package chance

import "math"

func (chance *Chance) Float() float64 {
	sign := 1
	if (chance.Bool()) {
		sign = -1
	}
	return math.MaxInt64 * chance.r.Float64() * sign
}

func (chance *Chance) FloatN(max float64) float64 {
	sign := 1
	if (chance.Bool()) {
		sign = -1
	}
	return chance.r.Float63n(max) * sign
}

func (chance *Chance) FloatBtw(min float64, max float64) float64 {
	return chance.r.Float64() * (max - min) + min
}

func (chance *Chance) Int() int {
	sign := 1
	if (chance.Bool()) {
		sign = -1
	}
	return chance.r.Int() * sign
}

func (chance *Chance) IntN(max int) int {
	return chance.r.Intn(max)
}
func (chance *Chance) IntBtw(min int, max int) int {
	return chance.r.Intn(max - min) + min
}

func (chance *Chance) Natural() int {
	return chance.r.Intn() + 1
}

func (chance *Chance) NaturalN(max int) int {
	return chance.r.Intn(max - 1) + 1
}


