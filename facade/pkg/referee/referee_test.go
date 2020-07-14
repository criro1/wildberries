package referee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReferee(t *testing.T) {
	r := NewReferee("Kassai", 0, 0)

	expectedYellow := "Kassai shows yellow card to Fernandez"
	t.Run("ShowYellowCard", func(t *testing.T) {
		resultYellow, err := r.ShowYellowCard("Fernandez")
		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expectedYellow, resultYellow)
	})

	expectedRed := "Kassai shows red card to James Rodriguez"
	t.Run("ShowRedCard", func(t *testing.T) {
		resultRed, err := r.ShowRedCard("James Rodriguez")
		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expectedRed, resultRed)
	})
	
	expected := "Referee Kassai shows 2 yellow cards and 1 red cards in this match"
	t.Run("ShowRedCard", func(t *testing.T) {
		r.ShowYellowCard("Valverde")
		result, err := r.GetStatistic()
		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expected, result)
	})
}
