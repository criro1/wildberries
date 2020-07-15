// testing of package visitor
package main

import (
	"fmt"

	vis "github.com/criro1/wildberries/visitor/pkg"
	serv "github.com/criro1/wildberries/visitor/pkg/services"
)

func main() {
	expect := "Moscow city buying:\nCustomer Petr bought pills paracetamol at the pharmacy `36'6`\nCustomer Petr bought goods cheese, sausages, pasta at the maket `Magnit`\nCustomer Petr got haircut side parted at the barbershop `Y Ludmili`"

	visitor := vis.NewCustomer("Petr")
	city := serv.NewCity("Moscow", "36'6", "Magnit", "Y Ludmili")

	result, err := city.DoPurchase(visitor, "paracetamol", "cheese, sausages, pasta", "side parted")
	if err != nil {
		fmt.Println("Error method DoPurchase")
		return
	}

	if result != expect {
		fmt.Println("Error, result != expect")
	} else {
		fmt.Println("Everything is OK! The result is:\n\n" + result)
	}
}
