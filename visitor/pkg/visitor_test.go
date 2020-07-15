package visitor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVisitPharmacy(t *testing.T) {
	c := NewCustomer("Marina")

	expected := "Customer Marina bought pills Analgin at the pharmacy `Stolichki`\n"
	t.Run("VisitPharmacy", func(t *testing.T) {
		result, err := c.VisitPharmacy("Stolichki", "Analgin")
		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expected, result)
	})
}

func TestVisitMarket(t *testing.T) {
	c := NewCustomer("Anna")

	expected := "Customer Anna bought goods chips, Pepsi at the maket `Pyaterochka`\n"
	t.Run("VisitMarket", func(t *testing.T) {
		result, err := c.VisitMarket("Pyaterochka", "chips, Pepsi")
		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expected, result)
	})
}

func TestVisitBarbershop(t *testing.T) {
	c := NewCustomer("Ivan Modnichkov")

	expected := "Customer Ivan Modnichkov got haircut lisii at the barbershop `PrichaBudetTop`"
	t.Run("VisitMarket", func(t *testing.T) {
		result, err := c.VisitBarbershop("PrichaBudetTop", "lisii")
		assert.NoError(t, err, "unexpected error:", err)
		assert.EqualValues(t, expected, result)
	})
}
