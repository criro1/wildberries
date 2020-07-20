// Package market ...
package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	unexpectedError = "unexpected error:"
	buyGoods        = "BuyGoods"
	masha           = "Masha"
	pyaterochka     = "Pyaterochka"
	expected        = "Customer Masha bought goods at the market `Pyaterochka`\n"
)

func TestAccept(t *testing.T) {
	t.Run(buyGoods, func(t *testing.T) {
		m := NewMarket(pyaterochka)
		result, err := m.BuyGoods(masha)
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expected, result)
	})
}
