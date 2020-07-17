// Package city ...
package city

import (
	mod "github.com/criro1/wildberries/visitor/pkg/models"
	"github.com/criro1/wildberries/visitor/pkg/city/pharmacy"
	"github.com/criro1/wildberries/visitor/pkg/city/market"
	"github.com/criro1/wildberries/visitor/pkg/city/barber"
	// vis "github.com/criro1/wildberries/visitor/pkg"
)

type visitor interface {
	VisitPharmacy(p *(pharmacy.Pharmacy)) (str string, err error)
	VisitMarket(m *(market.Market)) (str string, err error)
	VisitBarbershop(b *(barber.Barbershop)) (str string, err error)
}

type serv interface {
	Accept(v visitor) (str string, err error)
}

// City ...
type City interface {
	DoPurchase(v visitor) (str string, err error)
}

type city struct {
	name string
	serv []serv
}

// DoPurchase return string with buying from different services of the cuty
func (c *city) DoPurchase(v visitor) (str string, err error) {
	resulst := c.name + mod.CityBuy
	for _, place := range(c.serv) {
		s, err := place.Accept(v)
		if err != nil {
			return str, err
		}
		resulst += s
	}
	return resulst, nil
}

// NewCity ...
func NewCity(name string, phName, mktName, bbspName serv) City {
	return &city{
		name: name,
		serv: []serv{bbspName, mktName, phName},
	}
}
