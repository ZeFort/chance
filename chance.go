// Chance is a minimalist generator of random strings, numbers, etc. written in Go. Based on Chance.js
package chance

import (
	"math/rand"
	"time"
)

// Chance struct keeps own seed and random number generator
type Chance struct {
	seed int64
	r    *rand.Rand
}

// New creates new instance of Chance with seed based on current time
func New() *Chance {
	seed := time.Now().UTC().UnixNano()
	return &Chance{
		seed,
		rand.New(rand.NewSource(seed)),
	}
}

// NewS creates new instance of Chance with specified seed
func NewS(seed int64) *Chance {
	return &Chance{
		seed,
		rand.New(rand.NewSource(seed)),
	}
}
