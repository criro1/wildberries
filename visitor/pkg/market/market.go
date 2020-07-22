// Package market ...
package market

import (
	"errors"
	"strconv"

	"github.com/criro1/wildberries/visitor/pkg/api/v1"
)

type visitor interface {
	VisitMarket(m Market) (str string, err error)
}

// Market ...
type Market interface {
	Accept(v visitor) (str string, err error)
	ViewCameras() (str string, err error)
	GetName() (str string, err error)
}

type market struct {
	name string
	security int
}

// Accept accept the visitor
func (m *market) Accept(v visitor) (str string, err error) {
		return v.VisitMarket(m)
}

// GetName return name of the struct
func(m *market) GetName() (str string, err error) {
	if m.name == "" {
		err = errors.New(v1.BadMarkName)
		return
	}
	str = m.name
	return
}

// ViewCamereas return the string with information
func (m *market) ViewCameras() (str string, err error) {
	if m.security < 1 {
		err = errors.New(v1.BadMarkSec)
	}
	str = "One of " + strconv.Itoa(m.security) + " security guards of the market looked cameras"
	return
}

// NewMarket ...
func NewMarket(name string, security int) Market {
	return &market{
		name: name,
		security: security,
	}
}
