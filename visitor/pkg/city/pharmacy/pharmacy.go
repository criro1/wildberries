// Package pharmacy ...
package pharmacy

import (
	// "fmt"
	// "errors"

	// mod "github.com/criro1/wildberries/visitor/pkg/models"
	// "github.com/criro1/wildberries/visitor/pkg/city/market"
	// "github.com/criro1/wildberries/visitor/pkg/city/barber"
)

type visitor interface {
	VisitPharmacy(p *(Pharmacy)) (str string, err error)
	// VisitMarket(m *(market.Market)) (str string, err error)
	// VisitBarbershop(b *(barber.Barbershop)) (str string, err error)
}

// Phar ...
type Phar interface {
	Accept(v visitor) (str string, err error)
}

// Pharmacy ...
type Pharmacy struct {
	Name string
}

// Accept accept the visitor
func (p *Pharmacy) Accept(v visitor) (str string, err error) {
	return v.VisitPharmacy(p)
}

// NewPharmacy ...
func NewPharmacy(name string) Phar {
	return &Pharmacy{
		Name: name,
	}
}
