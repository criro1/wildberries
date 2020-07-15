// Package referee ...
package referee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	showYellowCard = "ShowYellowCard"
	expectedYellow = "Kassai shows yellow card to Fernandez"
	kassai = "Kassai"
	fernandez = "Fernandez"
	showRedCard = "ShowRedCard"
	expectedRed = "Kassai shows red card to James Rodriguez"
	james = "James Rodriguez"
	unexpError = "unexpected error:"
	expected = "Referee Kassai shows 2 yellow cards and 1 red cards in this match"
	statistic = "Statistic"
	valverde = "Valverde"
	lucas = "Lucas"
)

func TestRefereeYellow(t *testing.T) {
	t.Run(showYellowCard, func(t *testing.T) {
		r := NewReferee(kassai, 0, 0)
		resultYellow, err := r.ShowYellowCard(fernandez)
		assert.NoError(t, err, unexpError, err)
		assert.EqualValues(t, expectedYellow, resultYellow)
	})
}
func TestRefereeRed(t *testing.T) {
	t.Run(showRedCard, func(t *testing.T) {
		r := NewReferee(kassai, 0, 0)
		resultRed, err := r.ShowRedCard(james)
		assert.NoError(t, err, unexpError, err)
		assert.EqualValues(t, expectedRed, resultRed)
	})
}

func TestRefereeStat(t *testing.T) {
	t.Run(statistic, func(t *testing.T) {
		r := NewReferee(kassai, 0, 0)
		r.ShowYellowCard(valverde)
		r.ShowYellowCard(lucas)
		r.ShowRedCard(james)
		result, err := r.GetStatistic()
		assert.NoError(t, err, unexpError, err)
		assert.EqualValues(t, expected, result)
	})
}
