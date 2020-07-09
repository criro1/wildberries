// description of the package facade
package facade

import (
	"strings"
)

type interfaceFreekick interface {
	Todo() (string, error)
}

type freekick struct {
	footballers []InterfaceFootballer
	qty int
}

func (f *freekick) Todo() (string, error) {
	result := make([]string, f.qty, f.qty)
	for i, pl := range(f.footballers) {
		str, err := pl.Choose(i, f.qty)
		if err != nil {
			return "", err
		}
		result[i] = str
	}
	return strings.Join(result, "\n"), nil
}

// Newfreekick creates members of freekick
func NewFreekick(f... string) interfaceFreekick {
	l := len(f)
	footballers := make([]InterfaceFootballer, l, l)
	for i, pl := range(f) {
		footballers[i] = NewFootballer(pl)
	}

	return &freekick {
		footballers: footballers,
		qty: l,
	}	
}
