// testing of package facade
package main

import (
	"fmt"

	"github.com/criro1/wildberries/facade/pkg"
	"github.com/criro1/wildberries/facade/pkg/footballer"
	"github.com/criro1/wildberries/facade/pkg/referee"
)

func main() {
	expect := "1 - Ramos skips and don't touch the ball\n2 - Modric skips and don't touch the ball\n3 - Casemiro skips, but touchs and rolls the ball\n4 - Kroos kicks the ball\nReferee Scamina shows 2 yellow cards and 0 red cards in this match"

	players := footballer.NewFootballer("Ramos", "Modric", "Casemiro", "Kroos")
	referee := referee.NewReferee("Scamina", 0, 0)

	fk := facade.NewMatch(players, referee)
	result, err := fk.Todo("Mascelo", "Carvajal")
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
