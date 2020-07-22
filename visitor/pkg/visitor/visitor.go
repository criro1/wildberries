// Package visitor ...
package visitor

import (
	"fmt"

	"github.com/criro1/wildberries/visitor/pkg/api/v1"
	"github.com/criro1/wildberries/visitor/pkg/barbershop"
	"github.com/criro1/wildberries/visitor/pkg/market"
	"github.com/criro1/wildberries/visitor/pkg/pharmacy"
)

// Visitor ...
type Visitor interface {
	VisitPharmacy(p pharmacy.Pharmacy) (str string, err error)
	VisitMarket(m market.Market) (str string, err error)
	VisitBarbershop(b barbershop.Barbershop) (str string, err error)
	GetMood() (num int)
}

type customer struct {
	name string
	mood int
}

// GetMood return the mood from the string
func (c *customer) GetMood() (num int) {
	num = c.mood
	return
}

// VisitPharmacy return the string with the buying at the pharmacy
func (c *customer) VisitPharmacy(p pharmacy.Pharmacy) (str string, err error) {
	s := c.name + v1.WantedBuy + v1.AskedIf
	str, err = p.Masks()
	if err != nil {
		return
	}
	str = s + str
	return
}

// VisitMarket return the string with the buying at the market
func (c *customer) VisitMarket(m market.Market) (str string, err error) {
	cameras, err := m.ViewCameras()
	if err != nil {
		return
	}
	s := fmt.Sprintf(v1.LoseSmth, c.name, v1.Keys, m.GetName()) + cameras
	str = s + v1.DidntFind
	c.mood -= v1.MisusMood
	return
}

// VisitBarbershop return the string with the buying at the barbershop
func (c *customer) VisitBarbershop(b barbershop.Barbershop) (str string, err error) {
	s := c.name + v1.GreatHaircut
	c.mood += v1.PlusMood
	str1, err := b.SignUp(c.name, v1.Time)
	if err != nil {
		return
	}
	str2, err := b.BuyHaircut(c.name)
	if err != nil {
		return
	}
	str = s + str1 + str2 + v1.MoodBetter
	return
}

// NewCustomer ...
func NewCustomer(name string, mood int) Visitor {
	return &customer{
		name: name,
		mood: mood,
	}
}
