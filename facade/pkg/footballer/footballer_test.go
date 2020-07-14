package footballer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFootballerGet(t *testing.T) {
	f := NewFootballer("Varan", "Hazard", "Bale")

	expectedQty := 3
	t.Run("GetQty", func(t *testing.T) {
		resultQty, err := f.GetQty()
		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expectedQty, resultQty)
	})
}

func TestFootballerChoose(t *testing.T) {
	f := NewFootballer("Varan", "Hazard", "Bale")
	expectedChoose := "1 - Varan skips and don't touch the ball\n2 - Hazard skips, but touchs and rolls the ball\n3 - Bale kicks the ball\n"
	t.Run("Choose", func(t *testing.T) {
		resultQty, err := f.GetQty()
		resultChoose := ""
		for i := 0; i < 3; i++ {
			r, err := f.Choose(i, resultQty)
			if err != nil {
				t.Error("Error:", err)
			}
			resultChoose += r + "\n"
		}

		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expectedChoose, resultChoose)
	})
}
