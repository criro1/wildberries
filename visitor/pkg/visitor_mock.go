package visitor

import (
	"github.com/stretchr/testify/mock"
)

// MockVis ...
type MockVis struct {
	mock.Mock
}

// VisitPharmacy ...
func (m *MockVis) VisitPharmacy(pharmacy, pill string) (str string, err error) {
	args := m.Called()
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return str, args.Error(1)
}
// VisitMarket ...
func (m *MockVis) VisitMarket(pharmacy, pill string) (str string, err error) {
	args := m.Called()
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return str, args.Error(1)
}
// VisitBarbershop ...
func (m *MockVis) VisitBarbershop(pharmacy, pill string) (str string, err error) {
	args := m.Called()
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return str, args.Error(1)
}
