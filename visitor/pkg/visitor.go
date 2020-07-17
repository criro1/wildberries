// Package visitor ...
package visitor

import (
	"fmt"
	"errors"

	mod "github.com/criro1/wildberries/visitor/pkg/models"
	"github.com/criro1/wildberries/visitor/pkg/city/pharmacy"
	"github.com/criro1/wildberries/visitor/pkg/city/market"
	"github.com/criro1/wildberries/visitor/pkg/city/barber"
)

// Visitor ...
type Visitor interface {
	VisitPharmacy(p *(pharmacy.Pharmacy)) (str string, err error)
	VisitMarket(m *(market.Market)) (str string, err error)
	VisitBarbershop(b *(barber.Barbershop)) (str string, err error)
}

type customer struct {
	name string
}

// VisitPharmacy return the string with the buying at the pharmacy
func (c *customer) VisitPharmacy(p *(pharmacy.Pharmacy)) (str string, err error) {
	return p.BuyPills(c.name)
}

// BuyPills return the string with name of customer and market's name
func (p *pharmacy.Pharmacy) BuyPills(visName string) (str string, err error) {
	if p.Name == mod.EmptyStr {
		return str, errors.New(mod.BadPharName)
	}
	return fmt.Sprintf(mod.CustBuyPills, visName, p.Name), nil
}

// VisitMarket return the string with the buying at the market
func (c *customer) VisitMarket(m *(market.Market)) (str string, err error) {
	return m.BuyGoods(c.name)
}

// BuyGoods return the string with name of customer and market's name
func (m *market.Market) BuyGoods(visName string) (str string, err error) {
	if m.Name == mod.EmptyStr {
		return str, errors.New(mod.BadPharName)
	}
	return fmt.Sprintf(mod.CustByuGoods, visName, m.Name), nil
}

// VisitBarbershop return the string with the buying at the barbershop
func (c *customer) VisitBarbershop(b *(barber.Barbershop)) (str string, err error) {
	return b.BuyHaircut(c.name)
}

// BuyHaircut return the string with name of customer and haircut's name
func (m *barber.Barbershop) BuyHaircut(visName string) (str string, err error) {
	if m.Name == mod.EmptyStr {
		return str, errors.New(mod.BadPharName)
	}
	return fmt.Sprintf(mod.CustGetHaircut, visName, m.Name), nil
}

// NewCustomer ...
func NewCustomer(name string) Visitor {
	return &customer{
		name: name,
	}
}
