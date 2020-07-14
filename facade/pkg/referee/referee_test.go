package referee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReferee(t *testing.T) {
	r := NewReferee("Kassai", 0, 0)

	expectedYellow := "Kassai shows yellow card to Fernandez"
	resultYellow, err := r.ShowYellowCard("Fernandez")
	if err != nil {
		t.Error("Error:", err)
	}
	if expectedYellow != resultYellow {
		t.Error("Error, expectedYellow != resultYellow")
	}
	r.ShowYellowCard("Valverde")

	expectedRed := "Kassai shows red card to James Rodriguez"
	resultRed, err := r.ShowRedCard("James Rodriguez")
	if err != nil {
		t.Error("Error:", err)
	}
	if expectedRed != resultRed {
		t.Error("Error, expectedRed != resultRed")
	}

	expected := "Referee Kassai shows 2 yellow cards and 1 red cards in this match"
	result, err := r.GetStatistic()
	if err != nil {
		t.Error("Error:", err)
	}

	if expected != result {
		t.Error("Error, expected != result")
	}
}
