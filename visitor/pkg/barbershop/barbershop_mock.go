// Package barbershop ...
package barbershop

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

// SignUp ...
func (m *Mock) SignUp(customer string, time float64) (str string, err error) {
	args := m.Called(customer, time)
	return args.Get(0).(string), args.Error(1)
}

// BuyHaircut ...
func (m *Mock) BuyHaircut(visName string) (str string, err error) {
	args := m.Called(visName)
	return args.Get(0).(string), args.Error(1)
}
