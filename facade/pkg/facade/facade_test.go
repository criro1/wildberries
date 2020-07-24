// Package facade ...
package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"

	foot "github.com/criro1/wildberries/facade/pkg/footballer"
	ref "github.com/criro1/wildberries/facade/pkg/referee"
)

var (
	expected      = "1 - Pepe kicks the ball\nReferee Gusev shows 2 yellow cards and 1 red cards in this match\nFootballer Ibrahimovic got red card\n"
	add           = "Add"
	getQty        = "GetQty"
	choose        = "Choose"
	todo          = "Todo"
	pepeKick      = "1 - Pepe kicks the ball"
	showCard      = "ShowCard"
	ibrahimovic   = "Ibrahimovic"
	gusevShowY    = "Gusev shows yellow card to Ibrahimovic"
	gusevShowR    = "Gusev shows red card to Ibrahimovic"
	getStatistic  = "GetStatistic"
	refereeResult = "Referee Gusev shows 2 yellow cards and 1 red cards in this match\n"
	unexpErr      = "unexpected error:"
)

func TestFacade(t *testing.T) {
	t.Run(todo, func(t *testing.T) {
		footMock := new(foot.MockFb)

		footMock.On(add).Return(nil).Once()
		footMock.On(getQty).Return(1, nil).Once()
		footMock.On(choose, 0, 1).Return(pepeKick, nil).Once()

		refMock := new(ref.MockRef)

		refMock.On(showCard, ibrahimovic, true).Return(gusevShowY, nil).Twice()
		refMock.On(showCard, ibrahimovic, false).Return(gusevShowR, nil).Once()
		refMock.On(getStatistic).Return(refereeResult, nil).Once()

		match := NewMatch(
			footMock,
			refMock,
		)

		result, err := match.Todo(ibrahimovic, ibrahimovic)
		assert.NoError(t, err, unexpErr, err)
		assert.EqualValues(t, expected, result)
	})
}
