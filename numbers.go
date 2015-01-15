package chance

import (
	"math"
	"fmt"
	"strings"
)

func (chance *Chance) Float() float64 {
	sign := 1
	if (chance.Bool()) {
		sign = -1
	}
	return float64(math.MaxInt64) * chance.r.Float64() * float64(sign)
}

func (chance *Chance) FloatN(max float64) float64 {
	return chance.FloatBtw(0, max-1)
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
	return chance.r.Int() + 1
}

func (chance *Chance) NaturalN(max int) int {
	return chance.r.Intn(max - 1) + 1
}

func (chance *Chance) Phone(country string) string {
	templates := map[string]string{
		"BY" : "+375 29 ??? ????",
		"DE" : "+49 151 ????????",
		"US" : "+1 201-???-????",
		"RU" : "+7 912 ???-??-??",
		"CN" : " +86 131 ???? ????",
	}
	template := ""
	template = templates[country];
	if (template != ""){
		for i := 0; i < len(template); i++{
			if template[i] == '?' {
				symbol := chance.NumChar()
				template = strings.Replace(template, "?", string(symbol) , 1);
			}
		}
		return fmt.Sprintf("%s", template)
	}else{
		return fmt.Sprintf("%d-%d-%d", chance.IntBtw(100, 999), chance.IntBtw(10, 99), chance.IntBtw(10, 99))
	}

}


