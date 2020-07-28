// Package market ...
package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	unexpectedError = "unexpected error:"
	viewCameras     = "ViewCameras"
	pyaterochka     = "Pyaterochka"
	expectedCam     = "One of 5 security guards of the market looked cameras\n"
)

func TestCam(t *testing.T) {
	t.Run(viewCameras, func(t *testing.T) {
		m := NewMarket(pyaterochka, 5)
		result, err := m.ViewCameras()
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expectedCam, result)
	})
}
