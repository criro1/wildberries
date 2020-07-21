// Package facade ...
package facade

import (
	"fmt"
	"strings"
	"sync"
)

const (
	red = "Footballer %s got red card"
)
type footballer interface {
	Choose(i, qty int) (str string, err error)
	GetQty() (x int, err error)
	Add(p ...string) (err error)
}

type referee interface {
	ShowCard(player string, yellow bool) (str string, err error)
	GetStatistic() (str string, err error)
}

// MatchInt interface ...
type MatchInt interface {
	Todo(badGyus ...string) (str string, err error)
}

type match struct {
	footballers footballer
	referee     referee
}

// Todo return the string with all motions of the match
func (m *match) Todo(badGyus ...string) (str string, err error) {
	amount, err := m.footballers.GetQty()
	if err != nil {
		return
	}
	result := make([]string, amount + 1)
	for i := 0; i < amount; i++ {
		s, errNew := m.footballers.Choose(i, amount)
		if errNew != nil {
			err = errNew
			return
		}
		result[i] = s
	}

	mu := &sync.Mutex{}
	mp := make(map[string]int, len(badGyus))
	for _, bg := range badGyus {
		mu.Lock()
		if val, ok := mp[bg]; !ok {
			_, errNew := m.referee.ShowCard(bg, true)
			if errNew != nil {
				err = errNew
				return
			}
			mp[bg] = 1
		} else {
			if val == 1 {
				_, errNew := m.referee.ShowCard(bg, true)
				if errNew != nil {
					err = errNew
					return
				}
				_, errNew = m.referee.ShowCard(bg, false)
				if errNew != nil {
					err = errNew
					return
				}
				mp[bg] = 2
			}
		}
		mu.Unlock()
	}
	s, errNew := m.referee.GetStatistic()
	if errNew != nil {
		err = errNew
		return
	}
	str = strings.Join(result, "\n") + s + "\n"

	for key, val := range mp {
		if val == 2 {
			str += fmt.Sprintf(red, key) + "\n"
		}
	} 
	return
}

// NewMatch ...
func NewMatch(players footballer, referee referee) MatchInt {
	return &match{
		footballers: players,
		referee:     referee,
	}
}
