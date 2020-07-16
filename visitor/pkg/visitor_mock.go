// Package visitor ...
package visitor

import (
	"github.com/stretchr/testify/mock"
	serv "github.com/criro1/wildberries/visitor/pkg/services"
)

// MockVis ...
type MockVis struct {
	mock.Mock
}

// VisitPharmacy ...
func (m *MockVis) VisitPharmacy(p *(serv.Pharmacy)) (str string, err error) {
	args := m.Called(p)
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return str, args.Error(1)
}
// VisitMarket ...
func (m *MockVis) VisitMarket(mkt *(serv.Market)) (str string, err error) {
	args := m.Called(mkt)
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return str, args.Error(1)
}
// VisitBarbershop ...
func (m *MockVis) VisitBarbershop(b *(serv.Barbershop)) (str string, err error) {
	args := m.Called(b)
	if a, ok := args.Get(0).(string); ok {
		return a, args.Error(1)
	}
	return str, args.Error(1)
}
