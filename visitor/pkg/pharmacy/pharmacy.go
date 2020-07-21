// Package pharmacy ...
package pharmacy

import (
	"fmt"
	"errors"

	"github.com/criro1/wildberries/visitor/pkg/api/v1"
)

type visitor interface {
	VisitPharmacy(p Pharmacy) (str string, err error)
}

// Pharmacy ...
type Pharmacy interface {
	Accept(v visitor) (str string, err error)
	BuyPills(visName string) (str string, err error)
}

type pharmacy struct {
	name string
}

// Accept accept the visitor
func (p *pharmacy) Accept(v visitor) (str string, err error) {
	return v.VisitPharmacy(p)
}


// здесь надо добавить уникальностый метод и его юзать в VisitMarket


// BuyPills return the string with name of customer and market's name
func (p *pharmacy) BuyPills(visName string) (str string, err error) {
	if p.name == v1.EmptyStr {
		err = errors.New(v1.BadPharName)
		return
	}
	str = fmt.Sprintf(v1.CustBuyPills, visName, p.name)
	return
}

// NewPharmacy ...
func NewPharmacy(name string) Pharmacy {
	return &pharmacy{
		name: name,
	}
}
