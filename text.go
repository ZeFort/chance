package chance

import (
	"fmt"
)

func (chance *Chance) Syllable() string {
	if chance.Bool() {
		options := chance.AnyStringN("VC",2)
		if options == "VC" || options == "CC" {
			return fmt.Sprintf("%s%s", chance.VowelChar(), chance.ConsonantChar());
		}else{
			return fmt.Sprintf("%s%s", chance.ConsonantChar(), chance.VowelChar());
		}
	}else {
		return fmt.Sprintf("%s%s%s", chance.ConsonantChar(), chance.VowelChar(), chance.ConsonantChar());
	}
}

func (chance *Chance) Word() string{
	n := chance.IntBtw(2, 4);
	var str [4]string
	for i := 0; i < 4; i++ {
		str[i] = ""
	}
	for i := 0; i < n; i++ {
		str[i] = chance.Syllable()
	}
	return fmt.Sprintf("%s%s%s%s",str[0],str[1],str[2],str[3]);
}
