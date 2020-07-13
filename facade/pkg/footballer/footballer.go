// Package footballer ...
package footballer

import (
	"errors"
	"fmt"

	"github.com/criro1/wildberries/facade/pkg/models"
)

// Footballer interface ...
type Footballer interface {
	Choose(i, qty int) (string, error)
	GetQty() (int, error)
}

type footballer struct {
	name []string
	qty  int
}

// GetQty gets qty from struct footballer
func (f *footballer) GetQty() (int, error) {
	if f.qty < 0 {
		return 0, errors.New(models.BadAmount)
	}
	return f.qty, nil
}

// Choose chooses one of privat metods of footballer struct
func (f *footballer) Choose(i, qty int) (string, error) {
	if i < qty-2 {
		return f.skipWithoutTouch(i, f.name[i])
	} else if i == qty-2 {
		return f.skipAndTouch(i, f.name[i])
	} else if i == qty-1 {
		return f.kick(i, f.name[i])
	}
	return "", nil
}

func (f *footballer) kick(i int, name string) (string, error) {
	if i < 0 {
		return "", errors.New(models.BadAmount)
	}
	return fmt.Sprintf(models.KickBall, i+1, name), nil
}

func (f *footballer) skipAndTouch(i int, name string) (string, error) {
	if i < 0 {
		return "", errors.New(models.BadAmount)
	}
	return fmt.Sprintf(models.SkipAndTouchBall, i+1, name), nil
}

func (f *footballer) skipWithoutTouch(i int, name string) (string, error) {
	if i < 0 {
		return "", errors.New(models.BadAmount)
	}
	return fmt.Sprintf(models.WithoutTouch, i+1, name), nil
}

// NewFootballer ...
func NewFootballer(f ...string) Footballer {
	l := len(f)
	footballers := make([]string, l, l)
	for i, pl := range f {
		footballers[i] = pl
	}

	return &footballer{
		name: footballers,
		qty:  l,
	}
}
