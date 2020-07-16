// Package services ...
package services

import (
	"fmt"
	"errors"

	mod "github.com/criro1/wildberries/visitor/pkg/models"
)
// Market ...
type Market struct {
	Name string
}

func (m *Market) accept(v visitor) (str string, err error) {
	return v.VisitMarket(m)
}

// BuyGoods return the string with name of customer and market's name
func (m *Market) BuyGoods(visName string) (str string, err error) {
	if m.Name == mod.EmptyStr {
		return str, errors.New(mod.BadPharName)
	}
	return fmt.Sprintf(mod.CustByuGoods, visName, m.Name), nil
}
