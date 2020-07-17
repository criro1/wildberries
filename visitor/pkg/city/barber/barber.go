// Package barber ...
package barber

import (
	// "fmt"
	// "errors"

	// mod "github.com/criro1/wildberries/visitor/pkg/models"
	// "github.com/criro1/wildberries/visitor/pkg/city/pharmacy"
	// "github.com/criro1/wildberries/visitor/pkg/city/market"
	// vis "github.com/criro1/wildberries/visitor/pkg"
)

type visitor interface {
	// VisitPharmacy(p *(pharmacy.Pharmacy)) (str string, err error)
	// VisitMarket(m *(market.Market)) (str string, err error)
	VisitBarbershop(b *(Barbershop)) (str string, err error)
}

// Barber ...
type Barber interface {
	Accept(v visitor) (str string, err error)
}

// Barbershop ...
type Barbershop struct {
	Name string
}

// Accept accept the visitor
func (b *Barbershop) Accept(v visitor) (str string, err error) {
	return v.VisitBarbershop(b)
}

// NewBarbershop ...
func NewBarbershop(name string) Barber {
	return &Barbershop{
		Name: name,
	}
}
