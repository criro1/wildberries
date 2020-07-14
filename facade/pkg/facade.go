// Package facade ...
package facade

import (
	"strings"
)

type footballer interface {
	Choose(i, qty int) (string, error)
	GetQty() (int, error)
}

type referee interface {
	ShowYellowCard(player string) (string, error)
	ShowRedCard(player string) (string, error)
	GetStatistic() (string, error)
}

// MatchInt interface ...
type MatchInt interface {
	Todo(badGyus ...string) (string, error)
}

type match struct {
	footballers footballer
	referee     referee
}

// Todo return the string with all motions of the match
func (f *match) Todo(badGyus ...string) (str string, err error) {
	amount, err := f.footballers.GetQty()
	if err != nil {
		return str, err
	}
	result := make([]string, amount+1, amount+1)
	for i := 0; i < amount; i++ {
		str, err := f.footballers.Choose(i, amount)
		if err != nil {
			return str, err
		}
		result[i] = str
	}
	for _, bg := range badGyus {
		_, err := f.referee.ShowYellowCard(bg)
		if err != nil {
			return str, err
		}
	}
	result[amount], err = f.referee.GetStatistic()
	if err != nil {
		return str, err
	}
	return strings.Join(result, "\n"), nil
}

// NewMatch ...
func NewMatch(players footballer, referee referee) MatchInt {
	return &match{
		footballers: players,
		referee:     referee,
	}
}
