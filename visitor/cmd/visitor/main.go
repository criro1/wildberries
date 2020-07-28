// Package main ...
package main

import (
	"fmt"
	"strconv"

	"github.com/criro1/wildberries/visitor/pkg/barbershop"
	"github.com/criro1/wildberries/visitor/pkg/market"
	"github.com/criro1/wildberries/visitor/pkg/pharmacy"
	vis "github.com/criro1/wildberries/visitor/pkg/visitor"
)

const (
	expect        = "The mood was 78 y.e.\nMarket:\nPetr loosed his keys at the market Magnit\nOne of 3 security guards of the market looked cameras\nHe didn't find anything\nPharmacy:\nPetr wanted to buy masks\nHe asked, if the prahmacy had them\nUnfortunately we have no masks, come back later\nBarbershop:\nPetr wanted to make great haircut.He phone to barber\nAdministrator Alexandra signed up Petr at 18.40 o'clock\nCustomer Petr got haircut at the barbershop `Y Ludmili`\nHis mood became more better\nThe mood became 94 y.e.\n"
	petr          = "Petr"
	ph36_6        = "36'6"
	magnit        = "Magnit"
	yLudmili      = "Y Ludmili"
	alexandra     = "Alexandra"
	resNotExp     = "Error, result != expect"
	everythingOk  = "Everything is OK! The result is:\n\n"
	sevE          = 78
	three         = 3
	theMood       = "The mood was "
	theMoodBecome = "The mood became "
	yE            = " y.e.\n"
	marketStr     = "Market:\n"
	pharStr       = "Pharmacy:\n"
	barbStr       = "Barbershop:\n"
	falsePh       = false
)

func main() {
	v := vis.NewCustomer(petr, sevE)
	ph := pharmacy.NewPharmacy(ph36_6, falsePh)
	mk := market.NewMarket(magnit, three)
	bb := barbershop.NewBarbershop(yLudmili, alexandra)

	result := theMood + strconv.Itoa(v.GetMood()) + yE + marketStr
	s, err := mk.Accept(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	result += s

	result += pharStr
	s, err = ph.Accept(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	result += s

	result += barbStr
	s, err = bb.Accept(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	result += s + theMoodBecome + strconv.Itoa(v.GetMood()) + yE

	if result != expect {
		fmt.Println(resNotExp)
	} else {
		fmt.Print(everythingOk + result)
	}
}
