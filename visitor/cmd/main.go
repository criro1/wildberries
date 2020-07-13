// testing of package visitor
package main

import (
	"fmt"

	"github.com/criro1/wildberries/visitor/pkg"
	// "github.com/criro1/wildberries/visitor/pkg/footballer"
	// "github.com/criro1/wildberries/visitor/pkg/referee"
)

func main() {
	expect := ""

	// players := footballer.NewFootballer("Ramos", "Modric", "Casemiro", "Kroos")
	// referee := referee.NewReferee("Scamina", 0, 0)
	// fk := facade.NewMatch(players, referee)
	// result, err := fk.Todo("Mascelo", "Carvajal")
	if err != nil {
		fmt.Println("Error method Todo")
		return
	}

	if result != expect {
		fmt.Println("Error, result != expect")
	} else {
		fmt.Println("Everything is OK! The result is:\n\n" + result)
	}
}
