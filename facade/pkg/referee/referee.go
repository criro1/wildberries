// Package referee ...
package referee

import (
	"errors"
	"fmt"

	"github.com/criro1/wildberries/facade/pkg/api/v1"
)

// Referee interface ...
type Referee interface {
	ShowCard(player string, yellow bool) (str string, err error)
	GetStatistic() (str string, err error)
}

type referee struct {
	name       string
	yellowCard int
	redCard    int
}

// GetStatistic shows the statistic of the referee during the match
func (r *referee) GetStatistic() (str string, err error) {
	if r.yellowCard < 0 || r.redCard < 0 {
		err = errors.New(v1.ErrorReferee)
		return
	}
	str = fmt.Sprintf(v1.Statistic, r.name, r.yellowCard, r.redCard)
	return
}

// ShowCard add the amount of cards to referee and return the string with footballer, who was shown the card
func (r *referee) ShowCard(player string, yellow bool) (str string, err error) {
	if r.yellowCard < 0 || r.redCard < 0 {
		err = errors.New(v1.ErrorReferee)
		return
	}
	s := ""
	switch {
	case yellow == true:
		r.yellowCard++
		s = v1.YellowCard
	default:
		r.redCard++
		s = v1.RedCard
	}
	str = fmt.Sprintf(s, r.name, player)
	return
}

// NewReferee ...
func NewReferee(name string, yellow, red int) Referee {
	return &referee{
		name:       name,
		yellowCard: yellow,
		redCard:    red,
	}
}
