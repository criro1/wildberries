// Package footballer ...
package footballer

import (
	"github.com/stretchr/testify/mock"
)

// MockFb ...
type MockFb struct {
	mock.Mock
}

// Add ...
func (m *MockFb) Add(p ...string) (err error) {
	args := m.Called(p)
	return args.Error(0)
}

// GetQty ...
func (m *MockFb) GetQty() (x int, err error) {
	args := m.Called()
	return args.Get(0).(int), args.Error(1)
}

// Choose ...
func (m *MockFb) Choose(i, qty int) (str string, err error) {
	args := m.Called(i, qty)
	return args.Get(0).(string), args.Error(1)
}
