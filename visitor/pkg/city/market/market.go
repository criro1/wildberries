// Package market ...
package market

import (
	// "fmt"
	// "errors"

	// mod "github.com/criro1/wildberries/visitor/pkg/models"
	// "github.com/criro1/wildberries/visitor/pkg/city/pharmacy"
	// "github.com/criro1/wildberries/visitor/pkg/city/barber"
)

type visitor interface {
	// VisitPharmacy(p *(pharmacy.Pharmacy)) (str string, err error)
	VisitMarket(m *(Market)) (str string, err error)
	// VisitBarbershop(b *(barber.Barbershop)) (str string, err error)
}

// Mark ...
type Mark interface {
	Accept(v visitor) (str string, err error)
}

// Market ...
type Market struct {
	Name string
}
// Accept accept the visitor
func (m *Market) Accept(v visitor) (str string, err error) {
	return v.VisitMarket(m)
}

// NewMarket ...
func NewMarket(name string) Mark {
	return &Market{
		Name: name,
	}
}
