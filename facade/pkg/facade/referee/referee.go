// description of package footballer
package referee

import (
	"fmt"
	"errors"

	"github.com/criro1/wildberries/facade/pkg/facade/referee/models"
)

type Referee interface {
	ShowYellowCard(player string) (string, error)
	ShowRedCard(player string) (string, error)
	GetStatistic() (string, error)
}

type referee struct {
	name string
	yellowCard int
	redCard int
}

func (r *referee) GetStatistic() (string, error) {
	if r.yellowCard < 0 || r.redCard < 0 {
		return "", errors.New(models.ErrorReferee)
	}
	return fmt.Sprintf(models.Statistic, r.name, r.yellowCard, r.redCard), nil
}

func (r *referee) ShowYellowCard(player string) (string, error) {
	if r.yellowCard < 0 {
		return "", errors.New(models.ErrorReferee)
	}
	r.yellowCard++
	return fmt.Sprintf(models.YellowCard, r.name, player), nil
}

func (r *referee) ShowRedCard(player string) (string, error) {
	if r.redCard < 0 {
		return "", errors.New(models.ErrorReferee)
	}
	r.redCard++
	return fmt.Sprintf(models.RedCard, r.name, player), nil
}

// NewReferee ...
func NewReferee(name string, yellow, red int) Referee {
	return &referee{
		name: name,
		yellowCard: yellow,
		redCard: red,
	}
}
