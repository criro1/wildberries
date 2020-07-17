// Package main ...
package main

import (
	"fmt"

	vis "github.com/criro1/wildberries/visitor/pkg"
	mod "github.com/criro1/wildberries/visitor/pkg/models"
	pharmacy "github.com/criro1/wildberries/visitor/pkg/pharmacy"
	market "github.com/criro1/wildberries/visitor/pkg/market"
	barbershop "github.com/criro1/wildberries/visitor/pkg/barbershop"
)

const (
	expect       = "Moscow city buying:\nCustomer Petr got haircut at the barbershop `Y Ludmili`\nCustomer Petr bought goods at the maket `Magnit`\nCustomer Petr bought pills at the pharmacy `36'6`\n"
	petr         = "Petr"
	moscow       = "Moscow"
	ph36_6       = "36'6"
	magnit       = "Magnit"
	yLudmili     = "Y Ludmili"
	errorMethod  = "Error method DoPurchase"
	resNotExp    = "Error, result != expect"
	everythingOk = "Everything is OK! The result is:\n\n"
)

type visitor interface {
	VisitPharmacy(p pharmacy.Pharmacy) (str string, err error)
	VisitMarket(m market.Market) (str string, err error)
	VisitBarbershop(b barbershop.Barbershop) (str string, err error)
}

type places interface {
	Accept(v visitor) (str string, err error)
	Buy(visName string) (str string, err error)
}

func main() {
	v := vis.NewCustomer(petr)
	ph := pharmacy.NewPharmacy(ph36_6)
	mk := market.NewMarket(magnit)
	bb := barbershop.NewBarbershop(yLudmili)
	result := moscow + mod.CityBuy
	s, err := bb.Accept(v)
	if err != nil {
		fmt.Println("Error barbershop")
		return
	}
	result += s

	s, err = mk.Accept(v)
	if err != nil {
		fmt.Println("Error market")
		return
	}
	result += s

	s, err = ph.Accept(v)
	if err != nil {
		fmt.Println("Error pharmacy")
		return
	}
	result += s


	if result != expect {
		fmt.Println(resNotExp)
	} else {
		fmt.Print(everythingOk + result)
	}
}
