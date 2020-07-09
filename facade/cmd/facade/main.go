// testing of package facade
package main

import (
	"fmt"

	"github.com/criro1/wildberries/facade/pkg/facade"
)

func main() {
	expect := "1 - Ramos skips and don't touch the ball\n2 - Modric skips and don't touch the ball\n3 - Casemiro skips, but touchs and rolls the ball\n4 - Kroos kicks the ball"

	foorballers := facade.NewFreekick("Ramos", "Modric", "Casemiro", "Kroos")
	result, err := foorballers.Todo()
	if err != nil {
		fmt.Println("Error method Todo")
		return
	}

	if (result != expect) {
		fmt.Println("Error, result != expect")
	} else {
		fmt.Println("Everything is OK! The result is:\n\n" + result)
	}
}
