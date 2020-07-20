// Package main ...
package main

import (
	"fmt"

	vis "github.com/criro1/wildberries/visitor/pkg/visitor"
	pharmacy "github.com/criro1/wildberries/visitor/pkg/pharmacy"
	market "github.com/criro1/wildberries/visitor/pkg/market"
	barbershop "github.com/criro1/wildberries/visitor/pkg/barbershop"
)

const (
	expect       = "Moscow city buying:\nCustomer Petr got haircut at the barbershop `Y Ludmili`\nCustomer Petr bought goods at the market `Magnit`\nCustomer Petr bought pills at the pharmacy `36'6`\n"
	petr         = "Petr"
	moscow       = "Moscow"
	cityBuy      = " city buying:\n"
	ph36_6       = "36'6"
	magnit       = "Magnit"
	yLudmili     = "Y Ludmili"
	errorMethod  = "Error method DoPurchase"
	resNotExp    = "Error, result != expect"
	everythingOk = "Everything is OK! The result is:\n\n"
)

func main() {
	v := vis.NewCustomer(petr)
	ph := pharmacy.NewPharmacy(ph36_6)
	mk := market.NewMarket(magnit)
	bb := barbershop.NewBarbershop(yLudmili)
	result := moscow + cityBuy
	
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
