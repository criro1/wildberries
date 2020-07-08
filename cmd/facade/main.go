// testing of package facade
package main

import (
	"fmt"
	f "../../pkg/facade"
)

func main() {
	expect := "1 - Ramos skips and don't touch the ball\n2 - Modric skips and don't touch the ball\n3 - Casemiro skips, but touch and rolls the ball\n4 - Kroos kicks the ball"

	foorballers := f.Newfreekick("Ramos", "Modric", "Casemiro", "Kroos")
	result := foorballers.Todo()

	if (result != expect) {
		panic("Error, result != expect")
	} else {
		fmt.Println("Everything is OK! The result is:\n\n" + result)
	}
}
