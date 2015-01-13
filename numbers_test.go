package chance

import "testing"

func TestIntN( t *testing.T) {
	chance := NewS(1)
	if chance.IntN(1) != 0 {
		t.Fail()
	}
}
