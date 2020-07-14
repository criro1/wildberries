package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"
	foot "github.com/criro1/wildberries/facade/pkg/footballer"
	ref "github.com/criro1/wildberries/facade/pkg/referee"
)

func TestFacade(t *testing.T) {
	f := foot.NewFootballer("Vinisius", "Casemiro", "Modric", "Pepe")
	r := ref.NewReferee("Vorobyev", 0, 0)
	m := NewMatch(f, r)
	expected := "1 - Vinisius skips and don't touch the ball\n2 - Casemiro skips and don't touch the ball\n3 - Modric skips, but touchs and rolls the ball\n4 - Pepe kicks the ball\nReferee Vorobyev shows 3 yellow cards and 0 red cards in this match"
	t.Run("Todo", func(t *testing.T) {
		result, err := m.Todo("Ronaldo", "Messi", "Ibrahimovic")
		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expected, result)
	})

}