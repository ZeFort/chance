# chance
Chance is a minimalist generator of random strings, numbers, etc. written in **Go**. Based on [Chance.js](http://chancejs.com/)

[![Build Status](https://travis-ci.org/ZeFort/chance.svg)](https://travis-ci.org/ZeFort/chance) [![GoDoc](https://godoc.org/github.com/ZeFort/chance?status.svg)](https://godoc.org/github.com/ZeFort/chance) [![Coverage Status](https://coveralls.io/repos/ZeFort/chance/badge.png)](https://coveralls.io/r/ZeFort/chance)

#### Implemented functions
```go
type Chance
func New() *Chance
func NewS(seed int64) *Chance
func (chance *Chance) AlphaChar() string
func (chance *Chance) AnyChar(str string) string
func (chance *Chance) AnyString(str string) string
func (chance *Chance) AnyStringN(str string, len int) string
func (chance *Chance) Bool() bool
func (chance *Chance) BoolWithChance(likeliness int) bool
func (chance *Chance) Char() string
func (chance *Chance) ConsonantChar() string
func (chance *Chance) Float() float64
func (chance *Chance) FloatBtw(min float64, max float64) float64
func (chance *Chance) FloatN(max float64) float64
func (chance *Chance) Int() int
func (chance *Chance) IntBtw(min int, max int) int
func (chance *Chance) IntN(max int) int
func (chance *Chance) LowerChar() string
func (chance *Chance) Natural() int
func (chance *Chance) NaturalN(max int) int
func (chance *Chance) NumChar() string
func (chance *Chance) Phone(country string) string
func (chance *Chance) String() string
func (chance *Chance) StringN(len int) string
func (chance *Chance) Syllable() string
func (chance *Chance) UpperChar() string
func (chance *Chance) VowelChar() string
func (chance *Chance) Word() string
```

#### Notes
Documentation is available here: [godoc.org](https://godoc.org/github.com/ZeFort/chance).
Full information about code coverage is also available here: [chance on gocover.io](http://gocover.io/github.com/ZeFort/chance).

#### Support
If you do have a contribution for the package feel free to put up a Pull Request or open Issue.
