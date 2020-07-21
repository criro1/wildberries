// Package visitor ...
package visitor

import (
	"github.com/stretchr/testify/mock"
	pharmacy "github.com/criro1/wildberries/visitor/pkg/pharmacy"
	market "github.com/criro1/wildberries/visitor/pkg/market"
	barbershop "github.com/criro1/wildberries/visitor/pkg/barbershop"
)

// MockVis ...
type MockVis struct {
	mock.Mock
}

// VisitPharmacy ...
func (m *MockVis) VisitPharmacy(p pharmacy.Pharmacy) (str string, err error) {
	args := m.Called(p)
	return args.Get(0).(string), args.Error(1)
}

// VisitMarket ...
func (m *MockVis) VisitMarket(mkt market.Market) (str string, err error) {
	args := m.Called(mkt)
	return args.Get(0).(string), args.Error(1)
}

// VisitBarbershop ...
func (m *MockVis) VisitBarbershop(b barbershop.Barbershop) (str string, err error) {
	args := m.Called(b)
	return args.Get(0).(string), args.Error(1)
}
