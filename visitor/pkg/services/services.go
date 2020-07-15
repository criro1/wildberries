package services

import(
	// "fmt"

	mod "github.com/criro1/wildberries/visitor/pkg/models"
)

// Visitor interface ...
type Visitor interface {
	VisitPharmacy(pharmacy, pill string) (str string, err error)
	VisitMarket(market, goods string) (str string, err error)
	VisitBarbershop(barbershop, haircut string) (str string, err error)
}

// Service interface ...
type Service interface {
	SellTo(v Visitor, product string) (str string, err error)
}

// CityInterface ...
type CityInterface interface {
	DoPurchase(v Visitor, pill, goods, haircut string) (str string, err error)
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

func(c *city) DoPurchase(v Visitor, pill, goods, haircut string) (str string, err error) {
	s0, err := c.Serv[0].SellTo(v, pill)
	if err != nil {
		return str, err
	}
	s1, err := c.Serv[1].SellTo(v, goods)
	if err != nil {
		return str, err
	}
	s2, err := c.Serv[2].SellTo(v, haircut)
	if err != nil {
		return str, err
	}
	return c.name + mod.CityBuy + s0 + s1 + s2, nil
}

func(p *pharmacy) SellTo(v Visitor, product string) (str string, err error) {
	return v.VisitPharmacy(p.name, product)
}

func(m *market) SellTo(v Visitor, product string) (str string, err error) {
	return v.VisitMarket(m.name, product)
}

func(b *barbershop) SellTo(v Visitor, product string) (str string, err error) {
	return v.VisitBarbershop(b.name, product)
}

// NewCity ...
func NewCity(name, phName, mktName, bbspName string) CityInterface {
	// serv := make([]Service, 3, 3)
	// serv = append(serv, &pharmacy{phName}, &market{mktName}, &barbershop{bbspName})
	// fmt.Println(len(serv), cap(serv))
	return &city{
		name: name,
		Serv: []Service{&pharmacy{phName}, &market{mktName}, &barbershop{bbspName}},
	}
}
