// Package main ...
package main

import (
	"fmt"
	"strconv"

	vis "github.com/criro1/wildberries/visitor/pkg/visitor"
	pharmacy "github.com/criro1/wildberries/visitor/pkg/pharmacy"
	market "github.com/criro1/wildberries/visitor/pkg/market"
	barbershop "github.com/criro1/wildberries/visitor/pkg/barbershop"
)

const (
	expect		 = "The mood was 78 y.e.\nMarket:\nPetr loosed his keys at the market Magnit\nOne of 3 security guards of the market looked cameras\nHe didn't find anything\nPharmacy:\nPetr wanted to buy masks\nHe asked, if the prahmacy had them\nUnfortunately we have no masks, come back later\nBarbershop:\nPetr wanted to make great haircut.He phone to barber\nAdministrator Alexandra signed up Petr at 18.40 o'clock\nCustomer Petr got haircut at the barbershop `Y Ludmili`\nHis mood became more better\nThe mood became 94 y.e.\n"
	petr         = "Petr"
	ph36_6       = "36'6"
	magnit       = "Magnit"
	yLudmili     = "Y Ludmili"
	alexandra    = "Alexandra"
	resNotExp    = "Error, result != expect"
	everythingOk = "Everything is OK! The result is:\n\n"
)

func main() {
	v := vis.NewCustomer(petr, 78)
	ph := pharmacy.NewPharmacy(ph36_6, false)
	mk := market.NewMarket(magnit, 3)
	bb := barbershop.NewBarbershop(yLudmili, alexandra)
	
	num, err := v.GetMood()
	if err != nil {
		fmt.Println("Error mood:\n", err)
		return
	}
	result := "The mood was " + strconv.Itoa(num) + " y.e.\n"
	result += "Market:\n"
	s, err := mk.Accept(v)
	if err != nil {
		fmt.Println("Error market:\n", err)
		return
	}
	result += s

	result += "Pharmacy:\n"
	s, err = ph.Accept(v)
	if err != nil {
		fmt.Println("Error pharmacy:\n", err)
		return
	}
	result += s

	result += "Barbershop:\n"
	s, err = bb.Accept(v)
	if err != nil {
		fmt.Println("Error barbershop:\n", err)
		return
	}
	result += s

	num, err = v.GetMood()
	if err != nil {
		fmt.Println("Error mood:\n", err)
		return
	}
	result += "The mood became " + strconv.Itoa(num) + " y.e.\n"

	if result != expect {
		fmt.Println(resNotExp)
		fmt.Print(result)
	} else {
		fmt.Print(everythingOk + result)
	}
}
