// Package facade ...
package facade

import (
	"fmt"
	"strings"
	"sync"

	"github.com/criro1/wildberries/facade/pkg/api/v1"
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
	result := make([]string, amount+1)
	for i := 0; i < amount; i++ {
		str, err = m.footballers.Choose(i, amount)
		if err != nil {
			return
		}
		result[i] = str
	}

	mp, err := mapWork(m.referee, badGyus...)
	if err != nil {
		return
	}
	s, err := m.referee.GetStatistic()
	if err != nil {
		return
	}
	str = strings.Join(result, "\n") + s

	for key, val := range mp {
		if val == 2 {
			str += fmt.Sprintf(v1.Red, key)
		}
	}
	return
}

func mapWork(r referee, badGyus ...string) (mp map[string]int, err error) {
	mu := &sync.Mutex{}
	mp = make(map[string]int, len(badGyus))
	for _, bg := range badGyus {
		mu.Lock()
		val, ok := mp[bg]
		mu.Unlock()
		if !ok {
			_, err = r.ShowCard(bg, true)
			if err != nil {
				return
			}
			mu.Lock()
			mp[bg] = 1
			mu.Unlock()
		} else {
			if val == 1 {
				_, err = r.ShowCard(bg, true)
				if err != nil {
					return
				}
				_, err = r.ShowCard(bg, false)
				if err != nil {
					return
				}
				mu.Lock()
				mp[bg] = 2
				mu.Unlock()
			}
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
