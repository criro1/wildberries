// Package services ...
package services

import (
	mod "github.com/criro1/wildberries/visitor/pkg/models"
)

type visitor interface {
	VisitPharmacy(p *Pharmacy) (str string, err error)
	VisitMarket(m *Market) (str string, err error)
	VisitBarbershop(b *Barbershop) (str string, err error)
}

type serv interface {
	accept(v visitor) (str string, err error)
}

// CityInterface ...
type CityInterface interface {
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
		s, err := place.accept(v)
		if err != nil {
			return str, err
		}
		resulst += s
	}
	return resulst, nil
}

// NewCity ...
func NewCity(name, phName, mktName, bbspName string) CityInterface {
	return &city{
		name: name,
		serv: []serv{&Barbershop{bbspName}, &Market{mktName}, &Pharmacy{phName}},
	}
}
