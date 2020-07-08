// description of the package facade
package facade

import (
	"strings"
	"fmt"
)

type footballer struct {
	name string
}

type freekick struct {
	footballers []footballer
	qty int
}

func (p *freekick) Todo() string {
	result := make([]string, p.qty, p.qty)
	for i, pl := range(p.footballers) {
		if (i < p.qty - 2) {
			result[i] = pl.skip(i, false)
		} else if (i == p.qty - 2) {
			result[i] = pl.skip(i, true)
		} else {
			result[i] = pl.kick(i)
		}
		result[i] += " the ball"
	}
	return strings.Join(result, "\n")
}

func (f *footballer) kick(i int) string {
	return fmt.Sprintf("%d - %s kicks", i + 1, f.name)
}

func (f *footballer) skip(i int, roll bool) string {
	res := fmt.Sprintf("%d - %s skips", i + 1, f.name)
	if (roll == false) {
		return res + " and don't touch"
	} else {
		return res + ", but touch and rolls"
	}
	
}

// Newfreekick creates members of freekick
func Newfreekick(f... string) *freekick {
	l := len(f)
	footballers := make([]footballer, l, l)
	for i, pl := range(f) {
		footballers[i].name = pl
	}

	return &freekick {
		footballers: footballers,
		qty: l,
	}	
}
