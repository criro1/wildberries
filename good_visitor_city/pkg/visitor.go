// Package visitor ...
package visitor

import (
	"github.com/criro1/wildberries/visitor/pkg/city"
	// pharmacy "github.com/criro1/wildberries/visitor/pkg/city/pharmacy"
	// market "github.com/criro1/wildberries/visitor/pkg/city/market"
	// barbershop "github.com/criro1/wildberries/visitor/pkg/city/barbershop"
)

// Visitor ...
type Visitor interface {
	VisitPharmacy(p city.Pharmacy) (str string, err error)
	VisitMarket(m city.Market) (str string, err error)
	VisitBarbershop(b city.Barbershop) (str string, err error)
}

type customer struct {
	name string
}

// VisitPharmacy return the string with the buying at the pharmacy
func (c *customer) VisitPharmacy(p  city.Pharmacy) (str string, err error) {
	return p.Buy(c.name)
}

// VisitMarket return the string with the buying at the market
func (c *customer) VisitMarket(m city.Market) (str string, err error) {
	return m.Buy(c.name)
}

// VisitBarbershop return the string with the buying at the barbershop
func (c *customer) VisitBarbershop(b city.Barbershop) (str string, err error) {
	return b.Buy(c.name)
}

// NewCustomer ...
func NewCustomer(name string) Visitor {
	return &customer{
		name: name,
	}
}
