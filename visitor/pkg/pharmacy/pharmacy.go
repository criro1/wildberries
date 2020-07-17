package pharmacy

import (
	"fmt"
	"errors"

	mod "github.com/criro1/wildberries/visitor/pkg/models"
)

type visitor interface {
	VisitPharmacy(p Pharmacy) (str string, err error)
}

// Pharmacy ...
type Pharmacy interface {
	Accept(v visitor) (str string, err error)
	Buy(visName string) (str string, err error)
}

type pharmacy struct {
	name string
}

// Accept accept the visitor
func (p *pharmacy) Accept(v visitor) (str string, err error) {
	return v.VisitPharmacy(p)
}

// BuyPills return the string with name of customer and market's name
func (p *pharmacy) Buy(visName string) (str string, err error) {
	if p.name == mod.EmptyStr {
		return str, errors.New(mod.BadPharName)
	}
	return fmt.Sprintf(mod.CustBuyPills, visName, p.name), nil
}

// NewPharmacy ...
func NewPharmacy(name string) Pharmacy {
	return &pharmacy{
		name: name,
	}
}
