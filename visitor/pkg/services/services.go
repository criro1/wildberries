package services

import (
	mod "github.com/criro1/wildberries/visitor/pkg/models"
)

// Visitor interface ...
type Visitor interface {
	VisitPharmacy(pharmacy string) (str string, err error)
	VisitMarket(market string) (str string, err error)
	VisitBarbershop(barbershop string) (str string, err error)
}

// Service interface ...
type Service interface {
	sellTo(v Visitor) (str string, err error)
}

// CityInterface ...
type CityInterface interface {
	DoPurchase(v Visitor) (str string, err error)
}

type city struct {
	name string
	Serv []Service
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
func (c *city) DoPurchase(v Visitor) (str string, err error) {
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

func (p *pharmacy) sellTo(v Visitor) (str string, err error) {
	return v.VisitPharmacy(p.name)
}

func (m *market) sellTo(v Visitor) (str string, err error) {
	return v.VisitMarket(m.name)
}

func (b *barbershop) sellTo(v Visitor) (str string, err error) {
	return v.VisitBarbershop(b.name)
}

// NewCity ...
func NewCity(name, phName, mktName, bbspName string) CityInterface {
	return &city{
		name: name,
		Serv: []Service{&barbershop{bbspName}, &market{mktName}, &pharmacy{phName}},
	}
}
