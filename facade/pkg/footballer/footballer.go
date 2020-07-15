// Package footballer ...
package footballer

import (
	"fmt"
	"errors"

	"github.com/criro1/wildberries/facade/pkg/models"
)

// Footballer interface ...
type Footballer interface {
	Choose(i, qty int) (str string, err error)
	GetQty() (x int, err error)
}

type footballer struct {
	name []string
	qty  int
}

// GetQty gets qty from struct footballer
func (f *footballer) GetQty() (x int, err error) {
	if f.qty <= 0 {
		return x, errors.New(models.BadAmount)
	}
	return f.qty, nil
}

// Choose chooses one of privat metods of footballer struct
func (f *footballer) Choose(i, qty int) (str string, err error) {
	if i >= qty {
		return str, errors.New(models.BadAmount)
	}
	if i < qty-2 {
		return f.skipWithoutTouch(i, f.name[i])
	}
	if i == qty-2 {
		return f.skipAndTouch(i, f.name[i])
	}
	return f.kick(i, f.name[i])
}

func (f *footballer) kick(i int, name string) (str string, err error) {
	if i < 0 {
		return str, errors.New(models.BadAmount)
	}
	return fmt.Sprintf(models.KickBall, i + 1, name), nil
}

func (f *footballer) skipAndTouch(i int, name string) (str string, err error) {
	if i < 0 {
		return str, errors.New(models.BadAmount)
	}
	return fmt.Sprintf(models.SkipAndTouchBall, i + 1, name), nil
}

func (f *footballer) skipWithoutTouch(i int, name string) (str string, err error) {
	if i < 0 {
		return str, errors.New(models.BadAmount)
	}
	return fmt.Sprintf(models.WithoutTouch, i + 1, name), nil
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
