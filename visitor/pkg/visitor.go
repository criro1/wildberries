// Package visitor ...
package visitor

import (
	"fmt"
	"errors"

	mod "github.com/criro1/wildberries/visitor/pkg/models"
)

// CityInterface ...
type CityInterface interface {
}

// Services interface ...
type Services interface {
	SellTo(c *customer) string
}

// Visitor interface ...
type Visitor interface {
	VisitPharmacy(p *pharmacy, pill string) (str string, err error)
	VisitMarket(m *market, goods string) (str string, err error)
	VisitBarbershop(b *barbershop, haircut string) (str string, err error)
}

type customer struct {
	name string
}

type pharmacy struct {
	name string
}

type market struct {
	name string
}

type barbershop struct {
	name string
}

type city struct {
	name string
	serv []Services
}

func(p *pharmacy) SellTo(c customer) (str string, err error) {
	return c.VisitPharmacy(p, mod.Pill)
}

func(c *customer) VisitPharmacy(p *pharmacy, pill string) (str string, err error) {
	return p.buyPills(c.name, pill)
}

func(m *market) SellTo(c customer) (str string, err error) {
	return c.VisitMarket(m, mod.Goods)
}

func(c *customer) VisitMarket(m *market, goods string) (str string, err error) {
	return m.buyGoods(c.name, goods)
}

func(b *barbershop) SellTo(c customer) (str string, err error) {
	return c.VisitBarbershop(b, mod.Haircut)
}

func(c *customer) VisitBarbershop(b *barbershop, haircut string) (str string, err error) {
	return b.buyHaircut(c.name, haircut)
}

func(p *pharmacy) buyPills(cust, pill string) (str string, err error) {
	if p.name == "" {
		return str, errors.New(mod.BadPharName)
	}
	return fmt.Sprintf(mod.CustBuyPills, cust, pill, p.name), nil
}

func(m *market) buyGoods(cust, goods string) (str string, err error) {
	if m.name == "" {
		return str, errors.New(mod.BadMarkName)
	}
	return fmt.Sprintf(mod.CustByuGoods, cust, goods, m.name), nil
}

func(b *barbershop) buyHaircut(cust, haircut string) (str string, err error) {
	if b.name == "" {
		return str, errors.New(mod.BadBarbName)
	}
	return fmt.Sprintf(mod.CustGetHaircut, cust, haircut, b.name), nil
}

// NewCity ...
// func NewCity(name string) CityInterface {
// 	return &city{
// 		name: name,
// 		serv: []Services:[pharmacy, market, barbershop]
// 	}
// }

// NewCustomer ...
// func NewCustomer(name string) Visitor {
// 	return &customer{
// 		name: name,
// 	}
// }
