// Package referee ...
package referee

import (
	"fmt"
	"errors"

	"github.com/criro1/wildberries/facade/pkg/models"
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
		return str, errors.New(models.ErrorReferee)
	}
	return fmt.Sprintf(models.Statistic, r.name, r.yellowCard, r.redCard), nil
}

// ShowYellowCard add the amount of yellow cards to referee and return the string with footballer, who was shown the card
func (r *referee) ShowYellowCard(player string) (str string, err error) {
	if r.yellowCard < 0 {
		return str, errors.New(models.ErrorReferee)
	}
	r.yellowCard++
	return fmt.Sprintf(models.YellowCard, r.name, player), nil
}

// ShowRedCard add the amount of red cards to referee and return the string with footballer, who was shown the card
func (r *referee) ShowRedCard(player string) (str string, err error) {
	if r.redCard < 0 {
		return str, errors.New(models.ErrorReferee)
	}
	r.redCard++
	return fmt.Sprintf(models.RedCard, r.name, player), nil
}

// NewReferee ...
func NewReferee(name string, yellow, red int) Referee {
	return &referee{
		name:       name,
		yellowCard: yellow,
		redCard:    red,
	}
}
