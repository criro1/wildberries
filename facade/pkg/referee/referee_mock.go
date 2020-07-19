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

// ShowRedCard ...
func (m *MockRef) ShowRedCard(player string) (str string, err error) {
	args := m.Called(player)
	return args.Get(0).(string), args.Error(1)
}

// ShowYellowCard ...
func (m *MockRef) ShowYellowCard(player string) (str string, err error) {
	args := m.Called(player)
	return args.Get(0).(string), args.Error(1)
}
