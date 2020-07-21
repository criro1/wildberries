// Package visitor ...
package visitor

import (
	"errors"

	"github.com/criro1/wildberries/visitor/pkg/api/v1"
	pharmacy "github.com/criro1/wildberries/visitor/pkg/pharmacy"
	market "github.com/criro1/wildberries/visitor/pkg/market"
	barbershop "github.com/criro1/wildberries/visitor/pkg/barbershop"
)

// Visitor ...
type Visitor interface {
	VisitPharmacy(p pharmacy.Pharmacy) (str string, err error)
	VisitMarket(m market.Market) (str string, err error)
	VisitBarbershop(b barbershop.Barbershop) (str string, err error)
	GetMood() (num int, err error)
}

type customer struct {
	name string
	mood int
}

func (c *customer) GetMood() (num int, err error) {
	if c.mood < 0 {
		err = errors.New(v1.BadMood)
		return
	}
	num = c.mood
	return
}

// VisitPharmacy return the string with the buying at the pharmacy
func (c *customer) VisitPharmacy(p  pharmacy.Pharmacy) (str string, err error) {
	s := c.name + v1.WantedBuy
	s += v1.AskedIf
	strN, errN := p.Masks()
	if errN != nil {
		err = errN
		return
	}
	str = s + strN
	return
}

// VisitMarket return the string with the buying at the market
func (c *customer) VisitMarket(m market.Market) (str string, err error) {
	s := c.name + v1.NastySeller
	c.mood -= 20
	s += v1.BecauseOf
	str = s
	return
}

// VisitBarbershop return the string with the buying at the barbershop
func (c *customer) VisitBarbershop(b barbershop.Barbershop) (str string, err error) {
	s := c.name + v1.GreatHaircut
	strN, errN := b.BuyHaircut(c.name)
	if errN != nil {
		err = errN
		return
	}
	s += strN + v1.MoodBetter
	c.mood += 36
	str = s
	return
}

// NewCustomer ...
func NewCustomer(name string, mood int) Visitor {
	return &customer{
		name: name,
		mood: mood,
	}
}
