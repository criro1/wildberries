// Package services ...
package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	vis "github.com/criro1/wildberries/visitor/pkg"
)

const (
	buyHaircut      = "BuyHaircut"
	expected        = "Madrid city buying:\nCustomer Nickolay got haircut at the barbershop `Viktorya`\nCustomer Nickolay bought goods at the maket `Spar`\nCustomer Nickolay bought pills at the pharmacy `Y doma`\n"
	visitPharmacy   = "VisitPharmacy"
	visitMarket     = "VisitMarket"
	visitBarbershop = "VisitBarbershop"
	andrew          = "Andrew"
	madrid          = "Madrid"
	yDoma           = "Y doma"
	spar            = "Spar"
	viktorya        = "Viktorya"
	doPurchase      = "DoPurchase"
	retPharmacy     = "Customer Nickolay bought pills at the pharmacy `Y doma`\n"
	retMarker       = "Customer Nickolay bought goods at the maket `Spar`\n"
	retBarber       = "Customer Nickolay got haircut at the barbershop `Viktorya`\n"
	unexpectedError = "unexpected error:"
)

func TestDoPurchase(t *testing.T) {
	t.Run(doPurchase, func(t *testing.T) {
		visMock := new(vis.MockVis)

		visMock.On(visitPharmacy, &Pharmacy{Name: yDoma}).Return(retPharmacy, nil).Once()
		visMock.On(visitMarket, &Market{Name: spar}).Return(retMarker, nil).Once()
		visMock.On(visitBarbershop, &Barbershop{Name: viktorya}).Return(retBarber, nil).Once()
		
		c := NewCity(madrid, yDoma, spar, viktorya)
		result, err := c.DoPurchase(visMock)

		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expected, result)
	})
}

func TestBuyPills(t *testing.T) {
	t.Run(buyHaircut, func(t *testing.T) {
		b := &Pharmacy{Name: yDoma}
		res, err := b.BuyPills(andrew)

		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, "Customer Andrew bought pills at the pharmacy `Y doma`\n", res)
	})
}

func TestBuyGoods(t *testing.T) {
	t.Run(buyHaircut, func(t *testing.T) {
		b := &Market{Name: spar}
		res, err := b.BuyGoods(andrew)

		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, "Customer Andrew bought goods at the maket `Spar`\n", res)
	})
}

func TestBuyHaircut(t *testing.T) {
	t.Run(buyHaircut, func(t *testing.T) {
		b := &Barbershop{Name: viktorya}
		res, err := b.BuyHaircut(andrew)

		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, "Customer Andrew got haircut at the barbershop `Viktorya`\n", res)
	})
}
