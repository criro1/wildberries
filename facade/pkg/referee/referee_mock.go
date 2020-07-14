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
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return str, args.Error(1)
}

// ShowRedCard ...
func (m *MockRef) ShowRedCard(player string) (str string, err error) {
	args := m.Called(player)
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return str, args.Error(1)
}

// ShowYellowCard ...
func (m *MockRef) ShowYellowCard(player string) (str string, err error) {
	args := m.Called(player)
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return str, args.Error(1)
}
