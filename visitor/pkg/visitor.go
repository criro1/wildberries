package visitor

import (
	_"fmt"
)

// Visitor interface ...
type Visitor interface {
	VisitPharmacy(p *pharmacy) string
	VisitMarket(m *market) string
	VisitBarbershop(b *barbershop) string
}

// Services interface ...
type Services interface {}

type pharmacy struct{
	name string
}

type market struct{
	name string
}

type barbershop struct{
	name string
}
