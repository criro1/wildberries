// Package pharmacy ...
package pharmacy

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

// Masks ...
func (m *Mock) Masks() (str string, err error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}

// BuyPills ...
func (m *Mock) BuyPills(visName string) (str string, err error) {
	args := m.Called(visName)
	return args.Get(0).(string), args.Error(1)
}
