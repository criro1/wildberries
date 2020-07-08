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

func (p *freekick) Todo() (string, error) {
	result := make([]string, p.qty, p.qty)
	str := ""
	for i, pl := range(p.footballers) {
		if (i < p.qty - 2) {
			str, _ = pl.skipWithoutTouch(i)
		} else if (i == p.qty - 2) {
			str, _ = pl.skipAndTouch(i)
		} else {
			str, _ = pl.kick(i)
		}
		result[i] = str
	}
	return strings.Join(result, "\n"), nil
}

func (f *footballer) kick(i int) (string, error) {
	return fmt.Sprintf("%d - %s kicks the ball", i + 1, f.name), nil
}
func (f *footballer) skipAndTouch(i int) (string, error) {
	return fmt.Sprintf("%d - %s skips, but touchs and rolls the ball", i + 1, f.name), nil
}

func (f *footballer) skipWithoutTouch(i int) (string, error) {
	return fmt.Sprintf("%d - %s skips and don't touch the ball", i + 1, f.name), nil
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
