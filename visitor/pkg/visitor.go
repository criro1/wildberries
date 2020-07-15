// Package visitor ...
package visitor

import (
	"errors"
	"fmt"

	mod "github.com/criro1/wildberries/visitor/pkg/models"
)

// Visitor interface ...
type Visitor interface {
	VisitPharmacy(pharmacy string) (str string, err error)
	VisitMarket(market string) (str string, err error)
	VisitBarbershop(barbershop string) (str string, err error)
}

type customer struct {
	name string
}

// VisitPharmacy return the string with the buying at the pharmacy
func (c *customer) VisitPharmacy(pharmacy string) (str string, err error) {
	return c.buyPills(pharmacy)
}

// VisitMarket return the string with the buying at the market
func (c *customer) VisitMarket(market string) (str string, err error) {
	return c.buyGoods(market)
}

// VisitBarbershop return the string with the buying at the barbershop
func (c *customer) VisitBarbershop(barbershop string) (str string, err error) {
	return c.buyHaircut(barbershop)
}

func (c *customer) buyPills(pharmacy string) (str string, err error) {
	if pharmacy == mod.EmptyStr {
		return str, errors.New(mod.BadPharName)
	}
	return fmt.Sprintf(mod.CustBuyPills, c.name, pharmacy), nil
}

func (c *customer) buyGoods(market string) (str string, err error) {
	if market == mod.EmptyStr {
		return str, errors.New(mod.BadMarkName)
	}
	return fmt.Sprintf(mod.CustByuGoods, c.name, market), nil
}

func (c *customer) buyHaircut(barbershop string) (str string, err error) {
	if barbershop == mod.EmptyStr {
		return str, errors.New(mod.BadBarbName)
	}
	return fmt.Sprintf(mod.CustGetHaircut, c.name, barbershop), nil
}

// NewCustomer ...
func NewCustomer(name string) Visitor {
	return &customer{
		name: name,
	}
}
