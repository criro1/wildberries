// Package city ...
package city

import (
	"fmt"
	"errors"

	mod "github.com/criro1/wildberries/visitor/pkg/models"
)

type visitor interface {
	VisitBarbershop(b Barbershop) (str string, err error)
}

// Barbershop ...
type Barbershop interface {
	Accept(v visitor) (str string, err error)
	Buy(visName string) (str string, err error)
}

type barbershop struct {
	name string
}

// Accept accept the visitor
func (b *barbershop) Accept(v visitor) (str string, err error) {
	return v.VisitBarbershop(b)
}

// BuyHaircut return the string with name of customer and haircut's name
func (b *barbershop) Buy(visName string) (str string, err error) {
	if b.name == mod.EmptyStr {
		return str, errors.New(mod.BadPharName)
	}
	return fmt.Sprintf(mod.CustGetHaircut, visName, b.name), nil
}

// NewBarbershop ...
func NewBarbershop(name string) Barbershop {
	return &barbershop{
		name: name,
	}
}
