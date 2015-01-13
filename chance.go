package chance

import (
	"math/rand"
	"time"
)

type Chance struct {
	seed int64
	r Rand
}

func New() *Chance {
	seed := time.Now().UTC().UnixNano()
	return &Chance{
		seed,
		rand.New(seed),
	};
}

func New(seed int64) *Chance {
	return &Chance{
		seed,
		rand.New(seed),
	};
}
