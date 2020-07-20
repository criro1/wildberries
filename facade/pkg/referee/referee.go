// Package referee ...
package referee

import (
	"fmt"
	"errors"

	v1 "github.com/criro1/wildberries/facade/pkg/api/v1"
)

// Referee interface ...
type Referee interface {
	ShowYellowCard(player string) (str string, err error)
	ShowRedCard(player string) (str string, err error)
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

// ShowYellowCard add the amount of yellow cards to referee and return the string with footballer, who was shown the card
func (r *referee) ShowYellowCard(player string) (str string, err error) {
	if r.yellowCard < 0 {
		err = errors.New(v1.ErrorReferee)
		return
	}
	r.yellowCard++
	str = fmt.Sprintf(v1.YellowCard, r.name, player)
	return
}

// ShowRedCard add the amount of red cards to referee and return the string with footballer, who was shown the card
func (r *referee) ShowRedCard(player string) (str string, err error) {
	if r.redCard < 0 {
		err = errors.New(v1.ErrorReferee)
		return
	}
	r.redCard++
	str = fmt.Sprintf(v1.RedCard, r.name, player)
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
