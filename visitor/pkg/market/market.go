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
	GetName() (str string)
}

type market struct {
	name     string
	security int
}

// GetName return name of the struct
func (m *market) GetName() (str string) {
	str = m.name
	return
}

// Accept accept the visitor
func (m *market) Accept(v visitor) (str string, err error) {
	return v.VisitMarket(m)
}

// ViewCamereas return the string with information
func (m *market) ViewCameras() (str string, err error) {
	if m.security < v1.One {
		err = errors.New(v1.BadMarkSec)
	}
	str = v1.OneOf + strconv.Itoa(m.security) + v1.LookCam
	return
}

// NewMarket ...
func NewMarket(name string, security int) Market {
	return &market{
		name:     name,
		security: security,
	}
}
