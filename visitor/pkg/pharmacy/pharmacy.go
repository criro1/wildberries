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
	Masks() (str string, err error)
}

type pharmacy struct {
	name string
	masks bool
}

// Accept accept the visitor
func (p *pharmacy) Accept(v visitor) (str string, err error) {
	return v.VisitPharmacy(p)
}

func (p *pharmacy) Masks() (str string, err error) {
	if p.masks == true {
		str = v1.YesHave
	} else {
		str = v1.UnfDidnt
	}
	return
}

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
func NewPharmacy(name string, masks bool) Pharmacy {
	return &pharmacy{
		name: name,
		masks: masks,
	}
}
