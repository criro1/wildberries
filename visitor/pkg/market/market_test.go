// Package market ...
package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	unexpectedError = "unexpected error:"
	viewCameras		= "ViewCameras"
	masha           = "Masha"
	pyaterochka     = "Pyaterochka"
	expectedGetName	= "Pyaterochka"
	getName			= "GetName"
	expectedCam		= "One of 5 security guards of the market looked cameras"
)

func TestGetName(t *testing.T) {
	t.Run(getName, func(t *testing.T) {
		m := NewMarket(pyaterochka, 3)
		result, err := m.GetName()
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expectedGetName, result)
	})
}

func TestCam(t *testing.T) {
	t.Run(viewCameras, func(t *testing.T) {
		m := NewMarket(pyaterochka, 5)
		result, err := m.ViewCameras()
		assert.NoError(t, err, unexpectedError, err)
		assert.EqualValues(t, expectedCam, result)
	})
}
