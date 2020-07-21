// Package visitor ...
package visitor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/criro1/wildberries/visitor/pkg/pharmacy"
	"github.com/criro1/wildberries/visitor/pkg/market"
	"github.com/criro1/wildberries/visitor/pkg/barbershop"
)

const (
	unexpectedError = "unexpected error:"
	visitPharmacy   = "VisitPharmacy"
	andrey          = "Andrey"
	stolichki       = "Stolichki"
	expectedPh      = "Andrey wanted to buy masks\nHe asked, if the prahmacy had them\nYes, we have masks\n"
	visitMarket     = "VisitMarket"
	nick            = "Nick"
	pyaterochka     = "Pyaterochka"
	expectedMkt     = "Nick got nasty by the seller and his mood dropped\nBecause of this he did not buy anything\n"
	visitBarbershop = "VisitBarbershop"
	ivan            = "Ivan Modnichkov"
	prichaBudetTop  = "PrichaBudetTop"
	expectedBbshop  = "Ivan Modnichkov wanted to make great haircut\nCustomer Ivan Modnichkov got haircut at the barbershop `PrichaBudetTop`\nHis mood became more better\n"
	getMood         = "GetMood"
	alex            = "Alex"
	expGetMood		= 49
)

func TestGetMood(t *testing.T) {
t.Run(getMood, func(t *testing.T) {
	c := NewCustomer(alex, 49)
	resGetMood, err := c.GetMood()
	assert.NoError(t, err, unexpectedError, err)
	assert.EqualValues(t, expGetMood, resGetMood)
})
}

func TestVisitPharmacy(t *testing.T) {
	t.Run(visitPharmacy, func(t *testing.T) {
		c := NewCustomer(andrey, 55)
		resultPh, err := c.VisitPharmacy(pharmacy.NewPharmacy(stolichki, true))
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expectedPh, resultPh)
	})
}

func TestVisitMarket(t *testing.T) {
	t.Run(visitMarket, func(t *testing.T) {
		c := NewCustomer(nick, 80)
		resultMkt, err := c.VisitMarket(market.NewMarket(pyaterochka))
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expectedMkt, resultMkt)
	})
}

func TestVisitBarbershop(t *testing.T) {
	t.Run(visitBarbershop, func(t *testing.T) {
		c := NewCustomer(ivan, 73)
		resultBbshop, err := c.VisitBarbershop(barbershop.NewBarbershop(prichaBudetTop))
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expectedBbshop, resultBbshop)
	})
}
