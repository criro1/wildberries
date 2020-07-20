// Package footballer ...
package footballer

import (
	"fmt"
	"errors"

	v1 "github.com/criro1/wildberries/facade/pkg/api/v1"
)

// Footballer interface ...
type Footballer interface {
	Choose(i, qty int) (str string, err error)
	GetQty() (x int, err error)
	Add(p ...string) (err error)
}

type footballer struct {
	name []string
	qty  int
}

// Add new footballer into slice of their names
func (f *footballer) Add(p ...string) (err error) {
	if f.qty < 0 {
		err = errors.New(v1.BadAmount)
		return
	}
	f.name = append(f.name, p...)
	f.qty += len(p)
	return
}

// GetQty gets qty from struct footballer
func (f *footballer) GetQty() (x int, err error) {
	if f.qty < 0 {
		err = errors.New(v1.BadAmount)
		return
	}
	x = f.qty
	return
}

// Choose chooses one of privat metods of footballer struct
func (f *footballer) Choose(i, qty int) (str string, err error) {
	if i >= qty {
		err = errors.New(v1.BadAmount)
		return
	}
	if i < qty-2 {
		str, err = f.skipWithoutTouch(i, f.name[i])
		return
	}
	if i == qty-2 {
		str, err = f.skipAndTouch(i, f.name[i])
		return
	}
	str, err = f.kick(i, f.name[i])
	return
}

func (f *footballer) kick(i int, name string) (str string, err error) {
	if i < 0 {
		err = errors.New(v1.BadAmount)
		return
	}
	str = fmt.Sprintf(v1.KickBall, i + 1, name)
	return
}

func (f *footballer) skipAndTouch(i int, name string) (str string, err error) {
	if i < 0 {
		err = errors.New(v1.BadAmount)
		return
	}
	str = fmt.Sprintf(v1.SkipAndTouchBall, i + 1, name)
	return
}

func (f *footballer) skipWithoutTouch(i int, name string) (str string, err error) {
	if i < 0 {
		err = errors.New(v1.BadAmount)
		return
	}
	str = fmt.Sprintf(v1.WithoutTouch, i + 1, name)
	return
}

// NewFootballer ...
func NewFootballer() Footballer {
	return &footballer{
		name: []string{},
		qty:  0,
	}
}
