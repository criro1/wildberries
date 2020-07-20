// Package barbershop ...
package barbershop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	unexpectedError = "unexpected error:"
	expected        = "Customer Ivan Modnichkov got haircut at the barbershop `PrichaBudetTop`\n"
	buyHaircut      = "BuyHaircut"
	ivan            = "Ivan Modnichkov"
	prichaBudetTop  = "PrichaBudetTop"
)

func TestAccept(t *testing.T) {
	t.Run(buyHaircut, func(t *testing.T) {
		b := NewBarbershop(prichaBudetTop)
		result, err := b.BuyHaircut(ivan)
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expected, result)
	})
}
