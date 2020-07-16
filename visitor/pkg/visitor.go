// Package visitor ...
package visitor

import (
	serv "github.com/criro1/wildberries/visitor/pkg/services"
)

// Visitor ...
type Visitor interface {
	VisitPharmacy(p *(serv.Pharmacy)) (str string, err error)
	VisitMarket(m *(serv.Market)) (str string, err error)
	VisitBarbershop(b *(serv.Barbershop)) (str string, err error)
}

type customer struct {
	name string
}

// VisitPharmacy return the string with the buying at the pharmacy
func (c *customer) VisitPharmacy(p *(serv.Pharmacy)) (str string, err error) {
	return p.BuyPills(c.name)
}

// VisitMarket return the string with the buying at the market
func (c *customer) VisitMarket(m *(serv.Market)) (str string, err error) {
	return m.BuyGoods(c.name)
}

// VisitBarbershop return the string with the buying at the barbershop
func (c *customer) VisitBarbershop(b *(serv.Barbershop)) (str string, err error) {
	return b.BuyHaircut(c.name)
}

// NewCustomer ...
func NewCustomer(name string) Visitor {
	return &customer{
		name: name,
	}
}
