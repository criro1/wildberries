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
	admin			= "Anna"
	singUp 			= "SingUp"
	bob				= "Bob"
	time 			= 20.05
	expectedSingUp	= "Administrator Anna signed up Bob at 20.05 o'clock\n"
)

func TestSignUp(t *testing.T) {
	t.Run(singUp, func(t *testing.T) {
		b := NewBarbershop(prichaBudetTop, admin)
		result, err := b.SignUp(bob, time)
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expectedSingUp, result)
	})
}

func TestByu(t *testing.T) {
	t.Run(buyHaircut, func(t *testing.T) {
		b := NewBarbershop(prichaBudetTop, admin)
		result, err := b.BuyHaircut(ivan)
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expected, result)
	})
}
