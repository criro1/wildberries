// Package referee ...
package referee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	expectedYellow = "Kassai shows yellow card to Fernandez\n"
	kassai         = "Kassai"
	fernandez      = "Fernandez"
	showCard       = "ShowCard"
	expectedRed    = "Kassai shows red card to James Rodriguez\n"
	james          = "James Rodriguez"
	unexpError     = "unexpected error:"
	expected       = "Referee Kassai shows 2 yellow cards and 1 red cards in this match\n"
	statistic      = "Statistic"
	valverde       = "Valverde"
	lucas          = "Lucas"
)

func TestRefereeYellow(t *testing.T) {
	t.Run(showCard, func(t *testing.T) {
		r := NewReferee(kassai, 0, 0)
		resultYellow, err := r.ShowCard(fernandez, true)
		assert.NoError(t, err, unexpError, err)
		assert.EqualValues(t, expectedYellow, resultYellow)
	})
}
func TestRefereeRed(t *testing.T) {
	t.Run(showCard, func(t *testing.T) {
		r := NewReferee(kassai, 0, 0)
		resultRed, err := r.ShowCard(james, false)
		assert.NoError(t, err, unexpError, err)
		assert.EqualValues(t, expectedRed, resultRed)
	})
}

func TestRefereeStat(t *testing.T) {
	t.Run(statistic, func(t *testing.T) {
		r := NewReferee(kassai, 0, 0)
		r.ShowCard(valverde, true)
		r.ShowCard(lucas, true)
		r.ShowCard(james, false)
		result, err := r.GetStatistic()
		assert.NoError(t, err, unexpError, err)
		assert.EqualValues(t, expected, result)
	})
}
