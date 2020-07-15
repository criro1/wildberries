package visitor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	unexpectedError = "unexpected error:"
	visitPharmacy   = "VisitPharmacy"
	marina          = "Marina"
	stolichki       = "Stolichki"
	expectedPh      = "Customer Marina bought pills at the pharmacy `Stolichki`\n"
	visitMarket     = "VisitMarket"
	anna            = "Anna"
	pyaterochka     = "Pyaterochka"
	expectedMkt     = "Customer Anna bought goods at the maket `Pyaterochka`\n"
	visitBarbershop = "VisitBarbershop"
	ivan            = "Ivan Modnichkov"
	prichaBudetTop  = "PrichaBudetTop"
	expectedBbshop  = "Customer Ivan Modnichkov got haircut at the barbershop `PrichaBudetTop`\n"
)

func TestVisitPharmacy(t *testing.T) {
	t.Run(visitPharmacy, func(t *testing.T) {
		c := NewCustomer(marina)
		resultPh, err := c.VisitPharmacy(stolichki)
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expectedPh, resultPh)
	})
}

func TestVisitMarket(t *testing.T) {
	t.Run(visitMarket, func(t *testing.T) {
		c := NewCustomer(anna)
		resultMkt, err := c.VisitMarket(pyaterochka)
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expectedMkt, resultMkt)
	})
}

func TestVisitBarbershop(t *testing.T) {
	t.Run(visitBarbershop, func(t *testing.T) {
		c := NewCustomer(ivan)
		resultBbshop, err := c.VisitBarbershop(prichaBudetTop)
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expectedBbshop, resultBbshop)
	})
}
