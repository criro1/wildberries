package facade

import (
	"testing"

	foot "github.com/criro1/wildberries/facade/pkg/footballer"
	ref "github.com/criro1/wildberries/facade/pkg/referee"
	"github.com/stretchr/testify/assert"
)

var (
	expected       = "1 - Pepe kicks the ball\nReferee Gusev shows 1 yellow cards and 0 red cards in this match"
	GetQty         = "GetQty"
	Choose         = "Choose"
	PepeKick       = "1 - Pepe kicks the ball"
	ShowYellowCard = "ShowYellowCard"
	Ibrahimovic    = "Ibrahimovic"
	GusevShow      = "Gusev shows yellow card to Aguero"
	GetStatistic   = "GetStatistic"
	RefereeResult  = "Referee Gusev shows 1 yellow cards and 0 red cards in this match"
)

func TestFacade(t *testing.T) {
	t.Run("Todo", func(t *testing.T) {
		footMock := new(foot.MockFb)

		footMock.On(GetQty).Return(1, nil).Once()
		footMock.On(Choose, 0, 1).Return(PepeKick, nil).Once()

		refMock := new(ref.MockRef)

		refMock.On(ShowYellowCard, Ibrahimovic).Return(GusevShow, nil).Once()
		refMock.On(GetStatistic).Return(RefereeResult, nil).Once()

		match := NewMatch(
			footMock,
			refMock,
		)

		result, err := match.Todo(Ibrahimovic)
		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expected, result)
	})
}
