// Package footballer ...
package footballer

import (
	"github.com/stretchr/testify/mock"
)

// MockFb ...
type MockFb struct {
	mock.Mock
}

// GetQty ...
func (m *MockFb) GetQty() (x int, err error) {
	args := m.Called()
	if a, ok := args.Get(0).(int); ok {
		return a, args.Error(1)
	}
	return x, args.Error(1)
}

// Choose ...
func (m *MockFb) Choose(i, qty int) (str string, err error) {
	args := m.Called(i, qty)
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return str, args.Error(1)
}
