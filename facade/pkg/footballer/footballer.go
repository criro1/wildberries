// Package footballer ...
package footballer

import (
	"errors"
	"fmt"

	"github.com/criro1/wildberries/facade/pkg/api/v1"
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
	switch {
	case i >= qty:
		err = errors.New(v1.BadAmount)
	case i < qty-2:
		str, err = f.action(i, f.name[i], v1.WithoutTouch)
	case i == qty-2:
		str, err = f.action(i, f.name[i], v1.SkipAndTouchBall)
	default:
		str, err = f.action(i, f.name[i], v1.KickBall)
	}
	return
}

func (f *footballer) action(i int, name, act string) (str string, err error) {
	if i < 0 {
		err = errors.New(v1.BadAmount)
		return
	}
	str = fmt.Sprintf(act, i+1, name)
	return
}

// NewFootballer ...
func NewFootballer() Footballer {
	return &footballer{
		name: []string{},
		qty:  0,
	}
}
