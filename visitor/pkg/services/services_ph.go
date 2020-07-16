// Package services ...
package services

import (
	"fmt"
	"errors"

	mod "github.com/criro1/wildberries/visitor/pkg/models"
)

// Pharmacy ...
type Pharmacy struct {
	Name string
}

func (p *Pharmacy) accept(v visitor) (str string, err error) {
	return v.VisitPharmacy(p)
}

// BuyPills return the string with name of customer and market's name
func (p *Pharmacy) BuyPills(visName string) (str string, err error) {
	if p.Name == mod.EmptyStr {
		return str, errors.New(mod.BadPharName)
	}
	return fmt.Sprintf(mod.CustBuyPills, visName, p.Name), nil
}

// NewPharmacy ...
func NewPharmacy(name string) *Pharmacy {
	return &Pharmacy{
		Name: name,
	}
}
