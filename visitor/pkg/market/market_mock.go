// Package market ...
package market

import (
	"github.com/stretchr/testify/mock"
)

// Mock ...
type Mock struct {
	mock.Mock
}

// Accept ...
func (m *Mock) Accept(v visitor) (str string, err error) {
	args := m.Called(v)
	return args.Get(0).(string), args.Error(1)
}

// GetName ...
func (m *Mock) GetName() (str string) {
	args := m.Called()
	return args.Get(0).(string)
}

// ViewCameras ...
func (m *Mock) ViewCameras() (str string, err error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}
