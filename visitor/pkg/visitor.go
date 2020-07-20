// Package visitor ...
package visitor

import (
	pharmacy "github.com/criro1/wildberries/visitor/pkg/pharmacy"
	market "github.com/criro1/wildberries/visitor/pkg/market"
	barbershop "github.com/criro1/wildberries/visitor/pkg/barbershop"
)

// Visitor ...
type Visitor interface {
	VisitPharmacy(p pharmacy.Pharmacy) (str string, err error)
	VisitMarket(m market.Market) (str string, err error)
	VisitBarbershop(b barbershop.Barbershop) (str string, err error)
}

type customer struct {
	name string
}

// VisitPharmacy return the string with the buying at the pharmacy
func (c *customer) VisitPharmacy(p  pharmacy.Pharmacy) (str string, err error) {
	str, err = p.BuyPills(c.name)
	return
}

// VisitMarket return the string with the buying at the market
func (c *customer) VisitMarket(m market.Market) (str string, err error) {
	str, err = m.BuyGoods(c.name)
	return
}

// VisitBarbershop return the string with the buying at the barbershop
func (c *customer) VisitBarbershop(b barbershop.Barbershop) (str string, err error) {
	str, err = b.BuyHaircut(c.name)
	return
}

// NewCustomer ...
func NewCustomer(name string) Visitor {
	return &customer{
		name: name,
	}
}
