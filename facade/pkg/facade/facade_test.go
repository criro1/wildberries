// Package facade ...
package facade

import (
	"testing"

	foot "github.com/criro1/wildberries/facade/pkg/footballer"
	ref "github.com/criro1/wildberries/facade/pkg/referee"
	"github.com/stretchr/testify/assert"
)

var (
	expected       = "1 - Pepe kicks the ball\nReferee Gusev shows 1 yellow cards and 0 red cards in this match"
	add            = "Add"
	getQty         = "GetQty"
	choose         = "Choose"
	todo           = "Todo"
	pepeKick       = "1 - Pepe kicks the ball"
	showYellowCard = "ShowYellowCard"
	ibrahimovic    = "Ibrahimovic"
	gusevShow      = "Gusev shows yellow card to Ibrahimovic"
	getStatistic   = "GetStatistic"
	refereeResult  = "Referee Gusev shows 1 yellow cards and 0 red cards in this match"
	unexpErr       = "unexpected error:"
)

func TestFacade(t *testing.T) {
	t.Run(todo, func(t *testing.T) {
		footMock := new(foot.MockFb)

		footMock.On(add).Return(nil).Once()
		footMock.On(getQty).Return(1, nil).Once()
		footMock.On(choose, 0, 1).Return(pepeKick, nil).Once()

		refMock := new(ref.MockRef)

		refMock.On(showYellowCard, ibrahimovic).Return(gusevShow, nil).Once()
		refMock.On(getStatistic).Return(refereeResult, nil).Once()

		match := NewMatch(
			footMock,
			refMock,
		)

		result, err := match.Todo(ibrahimovic)
		assert.NoError(t, err, unexpErr, err)
		assert.EqualValues(t, expected, result)
	})
}
