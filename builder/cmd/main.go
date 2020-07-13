// testing of package builder
package main

import (
	"fmt"

	b "github.com/criro1/wildberries/builder/pkg"
)

func main() {
	expect := "The fish in rolls is tuna.\nThe vegetables in rolls are cucumber, asparagus.\nThe sause in rolls is wasabi."

	prod := new(b.Product)

	director := b.Director{&(b.RollsBuilder{prod})}
	director.Construct("tuna", "cucumber, asparagus", "wasabi")

	res, err := prod.GetIngridients()
	if err != nil {
		fmt.Println("Error")
	}
	if expect != res {
		fmt.Println("expect != res")
	} else {
		fmt.Println("Everything is OK, expect == result\n\n" + res)
	}

}
