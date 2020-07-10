// description of the package facade
package facade

import (
	// "fmt"
	"strings"
)

type Footballer interface {
	Choose(i, qty int) (string, error)
	GetQty() (int, error)
}

type Referee interface {
	ShowYellowCard(player string) (string, error)
	ShowRedCard(player string) (string, error)
	GetStatistic() (string, error)
}

type Freekick interface {
	Todo(badGyus... string) (string, error)
}

type freekick struct {
	footballers Footballer
	referee Referee
}

func (f *freekick) Todo(badGyus... string) (string, error) {
	amount, err := f.footballers.GetQty()
	if err != nil {
		return "", err
	}
	result := make([]string, amount + 1, amount + 1)
	for i := 0; i < amount; i++ {
		str, err := f.footballers.Choose(i, amount)
		if err != nil {
			return "", err
		}
		result[i] = str
	}
	for _, bg := range(badGyus) {
		_, err := f.referee.ShowYellowCard(bg)
		if err != nil {
			return "" ,err
		}
	}
	result[amount], err = f.referee.GetStatistic()
	if err != nil {
		return "", nil
	}
	return strings.Join(result, "\n"), nil
}

// Newfreekick creates members of freekick
func NewFreekick(players Footballer, referee Referee) Freekick {
	return &freekick {
		footballers: players,
		referee: referee,
	}	
}
