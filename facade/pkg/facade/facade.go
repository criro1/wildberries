// description of the package facade
package facade

import (
	"strings"
	"fmt"

	"github.com/criro1/wildberries/facade/pkg/facade/footballer"
)

type Freekick interface {
	Todo() (string, error)
}

type freekick struct {
	footballers []footballer
	qty int
}

func (f *freekick) Todo() (string, error) {
	result := make([]string, f.qty, f.qty)
	for i, pl := range(f.footballers) {
		str, err := pl.choose(i, f.qty)
		if err != nil {
			return "", err
		}
		result[i] = str
	}
	return strings.Join(result, "\n"), nil
}

// Newfreekick creates members of freekick
func NewFreekick(f... string) Freekick {
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
