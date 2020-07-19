// Package footballer ...
package footballer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	varan = "Varan"
	hazard = "Hazard"
	bale = "Bale"
	getQty = "GetQty"
	expectedQty = 3
	unexpError = "unexpected error:"
	expectedChoose = "1 - Varan skips and don't touch the ball\n2 - Hazard skips, but touchs and rolls the ball\n3 - Bale kicks the ball\n"
	choose = "Choose"
	add = "Add"
	errD = "Error:"
)

func TestFootballerAdd(t *testing.T) {
	t.Run(add, func(t *testing.T) {
		f := NewFootballer()
		err := f.Add(varan, hazard, bale)
		assert.NoError(t, err, unexpError, err)
		assert.EqualValues(t, "", "")
	})
}

func TestFootballerGet(t *testing.T) {
	t.Run(getQty, func(t *testing.T) {
		f := NewFootballer()
		f.Add(varan, hazard, bale)
		resultQty, err := f.GetQty()
		assert.NoError(t, err, unexpError, err)
		assert.EqualValues(t, expectedQty, resultQty)
	})
}

func TestFootballerChoose(t *testing.T) {
	t.Run(choose, func(t *testing.T) {
		f := NewFootballer()
		f.Add(varan, hazard, bale)
		resultQty, err := f.GetQty()
		resultChoose := ""
		for i := 0; i < 3; i++ {
			r, err := f.Choose(i, resultQty)
			if err != nil {
				t.Error(errD, err)
			}
			resultChoose += r + "\n"
		}

		assert.NoError(t, err, unexpError, err)
		assert.EqualValues(t, expectedChoose, resultChoose)
	})
}
