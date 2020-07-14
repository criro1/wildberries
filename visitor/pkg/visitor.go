// Package visitor ...
package visitor

import (
	"fmt"
	"errors"

	mod "github.com/criro1/wildberries/visitor/pkg/models"
)

// Visitor interface ...
type Visitor interface {
	VisitPharmacy(p *pharmacy) string
	VisitMarket(m *market) string
	VisitBarbershop(b *barbershop) string
}

// Customer struct ...
type Customer struct {
	name string
}

// Services interface ...
type Services interface {
	SellTo(c *Customer) string
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

// VisitPharmacy ...
func(c *Customer) VisitPharmacy(p *pharmacy, pill string) (string, error) {
	return p.BuyPills(c.name, pill)
}

func(p *pharmacy) BuyPills(cust, pill string) (string, error) {
	if p.name == "" {
		return "", errors.New("Bad pharmacy's name")
	}
	return fmt.Sprintf(mod.BuyPills, cust, pill, p.name), nil
}

func(c *Customer) VisitMarket(m *market, goods string) (string, error) {
	return m.BuyGoods(c.name, goods)
}

func(m *market) BuyGoods(cust, goods string) (string, error) {
	if p.name == "" {
		return "", errors.New("Bad market's name")
	}
	return fmt.Sprintf(mod.name, cust, pill, p.name), nil
}

func(c *Customer) VisitBarbershop(b *barbershop) string, error) {}
