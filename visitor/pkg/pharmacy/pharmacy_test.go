// Package pharmacy ...
package pharmacy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	unexpectedError = "unexpected error:"
	buyGoods        = "BuyGoods"
	anna            = "Anna"
	stolichki       = "Stolichki"
	expected        = "Customer Anna bought pills at the pharmacy `Stolichki`\n"
)

func TestAccept(t *testing.T) {
	t.Run(buyGoods, func(t *testing.T) {
		p := NewPharmacy(stolichki)
		result, err := p.BuyPills(anna)
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expected, result)
	})
}
