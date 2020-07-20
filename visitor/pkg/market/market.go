// Package market ...
package market

import (
	"fmt"
	"errors"

	"github.com/criro1/wildberries/visitor/pkg/api/v1"
)

type visitor interface {
	VisitMarket(m Market) (str string, err error)
}

// Market ...
type Market interface {
	Accept(v visitor) (str string, err error)
	BuyGoods(visName string) (str string, err error)
}

type market struct {
	name string
}
// Accept accept the visitor
func (m *market) Accept(v visitor) (str string, err error) {
	return v.VisitMarket(m)
}

// BuyGoods return the string with name of customer and market's name
func (m *market) BuyGoods(visName string) (str string, err error) {
	if m.name == v1.EmptyStr {
		err = errors.New(v1.BadPharName)
		return
	}
	str = fmt.Sprintf(v1.CustByuGoods, visName, m.name)
	return
}

// NewMarket ...
func NewMarket(name string) Market {
	return &market{
		name: name,
	}
}
