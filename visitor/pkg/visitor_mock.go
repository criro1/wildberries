package visitor

import (
	"github.com/stretchr/testify/mock"
)

// MockVis ...
type MockVis struct {
	mock.Mock
}

// VisitPharmacy ...
func (m *MockVis) VisitPharmacy(pharmacy string) (str string, err error) {
	args := m.Called(pharmacy)
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return str, args.Error(1)
}
// VisitMarket ...
func (m *MockVis) VisitMarket(market string) (str string, err error) {
	args := m.Called(market)
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return str, args.Error(1)
}
// VisitBarbershop ...
func (m *MockVis) VisitBarbershop(barber string) (str string, err error) {
	args := m.Called(barber)
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return str, args.Error(1)
}
