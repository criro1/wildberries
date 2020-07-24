// Package main ...
package main

import (
	"fmt"

	"github.com/criro1/wildberries/facade/pkg/facade"
	"github.com/criro1/wildberries/facade/pkg/footballer"
	"github.com/criro1/wildberries/facade/pkg/referee"
)

const (
	expect       = "1 - Ramos skips and don't touch the ball\n2 - Modric skips and don't touch the ball\n3 - Casemiro skips, but touchs and rolls the ball\n4 - Kroos kicks the ball\nReferee Scamina shows 3 yellow cards and 1 red cards in this match\nFootballer Mascelo got red card\n"
	resNotExtp   = "Error, result != expect"
	everythingOk = "Everything is OK! The result is:\n\n"
	errMethod    = "Error method Todo"
	ramos        = "Ramos"
	modric       = "Modric"
	casemiro     = "Casemiro"
	kroos        = "Kroos"
	scamina      = "Scamina"
	marselo      = "Mascelo"
	carvajal     = "Carvajal"
)

func main() {
	players := footballer.NewFootballer()
	err := players.Add(ramos, modric, casemiro, kroos)
	if err != nil {
		return
	}

	referee := referee.NewReferee(scamina, 0, 0)

	fk := facade.NewMatch(players, referee)
	result, err := fk.Todo(marselo, carvajal, marselo, marselo)
	if err != nil {
		fmt.Println(errMethod)
		return
	}

	if result != expect {
		fmt.Println(resNotExtp)
	} else {
		fmt.Print(everythingOk + result)
	}
}
