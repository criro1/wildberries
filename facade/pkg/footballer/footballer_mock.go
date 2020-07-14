package footballer

import (
	"github.com/stretchr/testify/mock"
)

// MockTest ...
type MockTest struct {
	mock.Mock
}

// GetQty ...
func (m *MockTest) GetQty() (int, error) {
	args := m.Called()
	if a, ok := args.Get(0).(int); ok {
		return a, args.Error(1)
	}
	return 0, args.Error(1)
}

// Choose ...
func (m *MockTest) Choose(i, qty int) (string, error) {
	args := m.Called(i, qty)
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return "", args.Error(1)
}
