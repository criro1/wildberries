package referee

import (
	"github.com/stretchr/testify/mock"
)

// MockFb ...
type MockFb struct {
	mock.Mock
}

// GetStatistic ...
func(m *MockFb) GetStatistic() (string, error) {
	args := m.Called()
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return "", args.Error(1)
}

// ShowRedCard ...
func(m *MockFb) ShowRedCard(player string) (string, error) {
	args := m.Called(player)
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return "", args.Error(1)
}

// ShowYellowCard ...
func(m *MockFb) ShowYellowCard(player string) (string, error) {
	args := m.Called(player)
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return "", args.Error(1)
}
