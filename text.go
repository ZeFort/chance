package chance

import (
	"fmt"
)

func (chance *Chance) Syllable(options string) string {
	if options == "VC" {
		return fmt.Sprintf("%s%s", chance.VowelChar() , chance.ConsonantChar());
	}
	if options == "CVC" {
		return fmt.Sprintf("%s%s%s", chance.ConsonantChar(), chance.VowelChar(), chance.ConsonantChar());
	}
	if options == "CV" {
		return fmt.Sprintf("%s%s", chance.ConsonantChar(), chance.VowelChar());
	}
	return fmt.Sprintf("%s%s", chance.ConsonantChar(), chance.VowelChar());
}

func (chance *Chance) Word() string{
	n := chance.IntBtw(2, 4);
	var str [4]string
	for i := 0; i < 4; i++ {
		str[i] = ""
	}
	for i := 0; i < n; i++ {
		if chance.Bool(){
			str[i] = chance.Syllable("VC")
		}else{
			str[i] = chance.Syllable("CVC")
		}
	}
	return fmt.Sprintf("%s%s%s%s",str[0],str[1],str[2],str[3]);
}
