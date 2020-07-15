// Package services ...
package services

import (
	mod "github.com/criro1/wildberries/visitor/pkg/models"
)

type visitor interface {
	VisitPharmacy(pharmacy string) (str string, err error)
	VisitMarket(market string) (str string, err error)
	VisitBarbershop(barbershop string) (str string, err error)
}

type service interface {
	sellTo(v visitor) (str string, err error)
}

// CityInterface ...
type CityInterface interface {
	DoPurchase(v visitor) (str string, err error)
}

type city struct {
	name string
	Serv []service
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

// DoPurchase return string with buying from different services of the cuty
func (c *city) DoPurchase(v visitor) (str string, err error) {
	resulst := c.name + mod.CityBuy
	for _, place := range(c.Serv) {
		s, err := place.sellTo(v)
		if err != nil {
			return str, err
		}
		resulst += s
	}
	return resulst, nil
}

func (p *pharmacy) sellTo(v visitor) (str string, err error) {
	return v.VisitPharmacy(p.name)
}

func (m *market) sellTo(v visitor) (str string, err error) {
	return v.VisitMarket(m.name)
}

func (b *barbershop) sellTo(v visitor) (str string, err error) {
	return v.VisitBarbershop(b.name)
}

// NewCity ...
func NewCity(name, phName, mktName, bbspName string) CityInterface {
	return &city{
		name: name,
		Serv: []service{&barbershop{bbspName}, &market{mktName}, &pharmacy{phName}},
	}
}
