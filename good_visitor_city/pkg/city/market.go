// Package city ...
package city

import (
	"fmt"
	"errors"

	mod "github.com/criro1/wildberries/visitor/pkg/models"
)

// Market ...
type Market interface {
	Accept(v visitor) (str string, err error)
	Buy(visName string) (str string, err error)
}

type market struct {
	name string
}
// Accept accept the visitor
func (m *market) Accept(v visitor) (str string, err error) {
	return v.VisitMarket(m)
}

// BuyGoods return the string with name of customer and market's name
func (m *market) Buy(visName string) (str string, err error) {
	if m.name == mod.EmptyStr {
		return str, errors.New(mod.BadPharName)
	}
	return fmt.Sprintf(mod.CustByuGoods, visName, m.name), nil
}

// NewMarket ...
func NewMarket(name string) Market {
	return &market{
		name: name,
	}
}
