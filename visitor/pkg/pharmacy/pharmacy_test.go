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
	masks           = "Masks"
	expectedTrue    = "Yes, we have masks\n"
	expectedFalse   = "Unfortunately we have no masks, come back later\n"
)

func TestMasksTrue(t *testing.T) {
	t.Run(masks, func(t *testing.T) {
		p := NewPharmacy(stolichki, true)
		result, err := p.Masks()
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expectedTrue, result)
	})
}

func TestMasksFalse(t *testing.T) {
	t.Run(masks, func(t *testing.T) {
		p := NewPharmacy(stolichki, false)
		result, err := p.Masks()
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expectedFalse, result)
	})
}

func TestByu(t *testing.T) {
	t.Run(buyGoods, func(t *testing.T) {
		p := NewPharmacy(stolichki, true)
		result, err := p.BuyPills(anna)
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expected, result)
	})
}
