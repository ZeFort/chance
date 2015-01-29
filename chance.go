package chance

import (
	"math/rand"
	"time"
)

type Chance struct {
	seed int64
	r    *rand.Rand
}

func New() *Chance {
	seed := time.Now().UTC().UnixNano()
	return &Chance{
		seed,
		rand.New(rand.NewSource(seed)),
	}
}

func NewS(seed int64) *Chance {
	return &Chance{
		seed,
		rand.New(rand.NewSource(seed)),
	}
}
