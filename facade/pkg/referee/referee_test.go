package referee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRefereeYellow(t *testing.T) {
	r := NewReferee("Kassai", 0, 0)

	expectedYellow := "Kassai shows yellow card to Fernandez"
	t.Run("ShowYellowCard", func(t *testing.T) {
		resultYellow, err := r.ShowYellowCard("Fernandez")
		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expectedYellow, resultYellow)
	})
}
func TestRefereeRed(t *testing.T) {
	r := NewReferee("Kassai", 0, 0)

	expectedRed := "Kassai shows red card to James Rodriguez"
	t.Run("ShowRedCard", func(t *testing.T) {
		resultRed, err := r.ShowRedCard("James Rodriguez")
		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expectedRed, resultRed)
	})
}

func TestRefereeStat(t *testing.T) {
	r := NewReferee("Kassai", 0, 0)

	expected := "Referee Kassai shows 2 yellow cards and 1 red cards in this match"
	t.Run("Statistic", func(t *testing.T) {
		r.ShowYellowCard("Valverde")
		r.ShowYellowCard("Lucas")
		r.ShowRedCard("James Rodriguez")
		result, err := r.GetStatistic()
		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expected, result)
	})
}
