// Package main ...
package main

import (
	"fmt"

	vis "github.com/criro1/wildberries/visitor/pkg"
	serv "github.com/criro1/wildberries/visitor/pkg/city"
	"github.com/criro1/wildberries/visitor/pkg/city/pharmacy"
	"github.com/criro1/wildberries/visitor/pkg/city/market"
	"github.com/criro1/wildberries/visitor/pkg/city/barber"
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

func main() {
	visitor := vis.NewCustomer(petr)
	city := serv.NewCity(moscow, pharmacy.NewPharmacy(ph36_6), market.NewMarket(magnit), barber.NewBarbershop(yLudmili))

	result, err := city.DoPurchase(visitor)
	if err != nil {
		fmt.Println(errorMethod)
		return
	}

	if result != expect {
		fmt.Println(resNotExp)
	} else {
		fmt.Print(everythingOk + result)
	}
}
