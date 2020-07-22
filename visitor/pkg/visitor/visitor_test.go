// Package visitor ...
package visitor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	
	"github.com/criro1/wildberries/visitor/pkg/barbershop"
	"github.com/criro1/wildberries/visitor/pkg/market"
	"github.com/criro1/wildberries/visitor/pkg/pharmacy"
)

const (
	unexpectedError = "unexpected error:"
	getMood         = "GetMood"
	visitPharmacy   = "VisitPharmacy"
	expectedPh      = "Andrey wanted to buy masks\nHe asked, if the prahmacy had them\nYes, we have masks\n"
	visitMarket     = "VisitMarket"
	expectedMkt     = "Nick loosed his keys at the market Pyaterochka\nOne of 6 security guards of the market looked cameras\nHe didn't find anything\n"
	visitBarbershop = "VisitBarbershop"
	expectedBbshop  = "Ivan Modnichkov wanted to make great haircut.He phone to barber\nAdministrator Alexandra signed up Petr at 18.40 o'clock\nCustomer Ivan Modnichkov got haircut at the barbershop `Y Ludmili`\nHis mood became more better\n"
	andrey          = "Andrey"
	ivan            = "Ivan Modnichkov"
	nick            = "Nick"
	vika            = "Victorya"
	pyaterochka     = "Pyaterochka"
	masks           = "Masks"
	mood1           = 55
	mood2           = 73
	mood3           = 5580
	yesHave         = "Yes, we have masks\n"
	viewCameras     = "ViewCameras"
	getName         = "GetName"
	oneOf           = "One of 6 security guards of the market looked cameras\n"
	signUp          = "SignUp"
	fl              = 18.4
	admRet          = "Administrator Alexandra signed up Petr at 18.40 o'clock\n"
	buyHaircut      = "BuyHaircut"
	hairRet         = "Customer Ivan Modnichkov got haircut at the barbershop `Y Ludmili`\n"
)

func TestVisitPharmacy(t *testing.T) {
	t.Run(visitPharmacy, func(t *testing.T) {
		p := new(pharmacy.Mock)
		p.On(masks).Return(yesHave, nil).Once()
		c := NewCustomer(andrey, mood1)
		resultPh, err := c.VisitPharmacy(p)
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expectedPh, resultPh)
	})
}

func TestVisitMarket(t *testing.T) {
	t.Run(visitMarket, func(t *testing.T) {
		m := new(market.Mock)
		m.On(getName).Return(pyaterochka).Once()
		m.On(viewCameras).Return(oneOf, nil).Once()
		c := NewCustomer(nick, mood2)
		resultMkt, err := c.VisitMarket(m)
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expectedMkt, resultMkt)
	})
}

func TestVisitBarbershop(t *testing.T) {
	t.Run(visitBarbershop, func(t *testing.T) {
		b := new(barbershop.Mock)
		b.On(signUp, ivan, fl).Return(admRet, nil).Once()
		b.On(buyHaircut, ivan).Return(hairRet, nil).Once()
		c := NewCustomer(ivan, mood3)
		resultBbshop, err := c.VisitBarbershop(b)
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expectedBbshop, resultBbshop)
	})
}
