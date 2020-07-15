// Package visitor ...
package visitor

import (
	"fmt"
	"errors"

	mod "github.com/criro1/wildberries/visitor/pkg/models"
)

// Service interface ...
type Service interface {
	SellTo(c *customer) (str string)
}

// Visitor interface ...
type Visitor interface {
	VisitPharmacy(pharmacy, pill string) (str string, err error)
	VisitMarket(market, goods string) (str string, err error)
	VisitBarbershop(barbershop, haircut string) (str string, err error)
}

type customer struct {
	name string
}

func(c *customer) VisitPharmacy(pharmacy, pill string) (str string, err error) {
	return c.buyPills(pharmacy, pill)
}

func(c *customer) VisitMarket(market, goods string) (str string, err error) {
	return c.buyGoods(market, goods)
}

func(c *customer) VisitBarbershop(barbershop, haircut string) (str string, err error) {
	return c.buyHaircut(barbershop, haircut)
}

func(c *customer) buyPills(pharmacy, pill string) (str string, err error) {
	if pharmacy == "" {
		return str, errors.New(mod.BadPharName)
	}
	return fmt.Sprintf(mod.CustBuyPills, c.name, pill, pharmacy), nil
}

func(c *customer) buyGoods(market, goods string) (str string, err error) {
	if market == "" {
		return str, errors.New(mod.BadMarkName)
	}
	return fmt.Sprintf(mod.CustByuGoods, c.name, goods, market), nil
}

func(c *customer) buyHaircut(barbershop, haircut string) (str string, err error) {
	if barbershop == "" {
		return str, errors.New(mod.BadBarbName)
	}
	return fmt.Sprintf(mod.CustGetHaircut, c.name, haircut, barbershop), nil
}

// NewCustomer ...
func NewCustomer(name string) Visitor {
	return &customer{
		name: name,
	}
}
