package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	vis "github.com/criro1/wildberries/visitor/pkg"
)

var (
	expected = "Madrid city buying:\nCustomer Nickolay bought pills Ibuprofen at the pharmacy `Y doma`\nCustomer Nickolay bought goods shampoo, bread, beer at the maket `Spar`\nCustomer Nickolay got haircut polubox at the barbershop `Viktorya`"
	VisitPharmacy = "VisitPharmacy"
	VisitMarket = "VisitMarket"
	VisitBarbershop = "VisitBarbershop"
	Madrid = "Madrid"
	Ydoma = "Y doma"
	Spar = "Spar"
	Viktorya = "Viktorya"
	DoPurchase = "DoPurchase"
	Ibuprofen = "Ibuprofen"
	RetPharmacy = "Customer Nickolay bought pills Ibuprofen at the pharmacy `Y doma`\n"
	Goods = "shampoo, bread, beer"
	RetMarker = "Customer Nickolay bought goods shampoo, bread, beer at the maket `Spar`\n"
	Polubox = "polubox"
	RetBarber = "Customer Nickolay got haircut polubox at the barbershop `Viktorya`"
)

func TestDoPurchase(t *testing.T) {
	t.Run(DoPurchase, func(t *testing.T) {
		c := NewCity(Madrid, Ydoma, Spar, Viktorya)
		visMock := new(vis.MockVis)
		visMock.On(VisitPharmacy, Ydoma, Ibuprofen).Return(RetPharmacy, nil).Once()
		visMock.On(VisitMarket, Spar, Goods).Return(RetMarker, nil).Once()
		visMock.On(VisitBarbershop, Viktorya, Polubox).Return(RetBarber, nil).Once()
		
		result, err := c.DoPurchase(visMock, Ibuprofen, Goods, Polubox)

		// v := vis.NewCustomer("Nickolay")
		// result, err := c.DoPurchase(v, Ibuprofen, Goods, Polubox)

		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expected, result)
	})
}
