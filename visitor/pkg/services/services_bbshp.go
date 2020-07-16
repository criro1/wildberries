// Package services ...
package services

import (
	"fmt"
	"errors"

	mod "github.com/criro1/wildberries/visitor/pkg/models"
)

// Barbershop ...
type Barbershop struct {
	Name string
}

func (b *Barbershop) accept(v visitor) (str string, err error) {
	return v.VisitBarbershop(b)
}

// BuyHaircut return the string with name of customer and haircut's name
func (b *Barbershop) BuyHaircut(visName string) (str string, err error) {
	if b.Name == mod.EmptyStr {
		return str, errors.New(mod.BadPharName)
	}
	return fmt.Sprintf(mod.CustGetHaircut, visName, b.Name), nil
}
