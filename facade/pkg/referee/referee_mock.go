// Package referee ...
package referee

import (
	"github.com/stretchr/testify/mock"
)

// MockRef ...
type MockRef struct {
	mock.Mock
}

// GetStatistic ...
func (m *MockRef) GetStatistic() (str string, err error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}

// ShowCard ...
func (m *MockRef) ShowCard(player string, yellow bool) (str string, err error) {
	args := m.Called(player, yellow)
	return args.Get(0).(string), args.Error(1)
}
