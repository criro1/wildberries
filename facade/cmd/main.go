// Package main ...
package main

import (
	"fmt"

	"github.com/criro1/wildberries/facade/pkg"
	"github.com/criro1/wildberries/facade/pkg/footballer"
	"github.com/criro1/wildberries/facade/pkg/referee"
)

const (
	expect = "1 - Ramos skips and don't touch the ball\n2 - Modric skips and don't touch the ball\n3 - Casemiro skips, but touchs and rolls the ball\n4 - Kroos kicks the ball\nReferee Scamina shows 2 yellow cards and 0 red cards in this match"
	resNotExtp = "Error, result != expect"
	everythingOk = "Everything is OK! The result is:\n\n"
	errMethod = "Error method Todo"
	ramos = "Ramos"
	modric = "Modric"
	casemiro = "Casemiro"
	kroos = "Kroos"
	scamina = "Scamina"
	marselo = "Mascelo"
	carvajal = "Carvajal"
)

func main() {
	players := footballer.NewFootballer(ramos, modric, casemiro, kroos)
	referee := referee.NewReferee(scamina, 0, 0)
	fk := facade.NewMatch(players, referee)
	result, err := fk.Todo(marselo, carvajal)
	if err != nil {
		fmt.Println(errMethod)
		return
	}

	if result != expect {
		fmt.Println(resNotExtp)
	} else {
		fmt.Println(everythingOk + result)
	}
}
