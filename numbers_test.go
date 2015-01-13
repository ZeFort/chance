package chance

import "testing"

func TestRInt( t *testing.T) {
	chance := NewS(1)
	if chance.RInt() != 0 {
		t.Fail()
	}
}
