// Package barbershop ...
package barbershop

import (
	"errors"
	"fmt"

	"github.com/criro1/wildberries/visitor/pkg/api/v1"
)

type visitor interface {
	VisitBarbershop(b Barbershop) (str string, err error)
}

// Barbershop ...
type Barbershop interface {
	Accept(v visitor) (str string, err error)
	BuyHaircut(visName string) (str string, err error)
	SignUp(customer string, time float64) (str string, err error)
}

type barbershop struct {
	name  string
	admin string
}

// Accept accept the visitor
func (b *barbershop) Accept(v visitor) (str string, err error) {
	return v.VisitBarbershop(b)
}

// SignUp signed up the customer for the haircut on time
// 0, 20, 9, 0.59 won't be changed, it's working time of the barber
func (b *barbershop) SignUp(customer string, time float64) (str string, err error) {
	if time < 0 || int(time) > 20 || int(time) < 9 || time-float64(int(time)) > 0.59 {
		err = errors.New(v1.BadTime)
		return
	}
	str = fmt.Sprintf(v1.AdmSignUp, b.admin, customer, time)
	return
}

// BuyHaircut return the string with name of customer and haircut's name
func (b *barbershop) BuyHaircut(visName string) (str string, err error) {
	if b.name == "" {
		err = errors.New(v1.BadPharName)
		return
	}
	str = fmt.Sprintf(v1.CustGetHaircut, visName, b.name)
	return
}

// NewBarbershop ...
func NewBarbershop(name, admin string) Barbershop {
	return &barbershop{
		name:  name,
		admin: admin,
	}
}
