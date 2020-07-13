package builder

import (
	_ "fmt"
	"errors"
	"strings"

	m "github.com/criro1/wildberries/builder/pkg/builder/models"
)


type Builder interface {
	MakeFish(str string) error
	MakeVeg(str string) error
	MakeSauce(str string) error
}

type Product struct {
	Ingredients string
}

type RollsBuilder struct {
	Product *Product
}

type Director struct {
	Builder Builder
}

func(d *Director) Construct(fish, veg, sause string) {
	d.Builder.MakeFish(fish)
	d.Builder.MakeVeg(veg)
	d.Builder.MakeSauce(sause)
}

func(r *RollsBuilder) MakeFish(fish string) error {
	if fish == "" {
		return errors.New(m.Wrong + m.Fish)
	}
	r.Product.Ingredients += m.FishInRolls + fish + m.Point
	return nil
}

func(r *RollsBuilder) MakeVeg(veg string) error {
	if veg == "" {
		return errors.New(m.Wrong + m.Veg)
	}
	if strings.Contains(veg, ",") {
		r.Product.Ingredients += m.VegAre + veg + m.Point
	} else {
		r.Product.Ingredients += m.VegIs + veg + m.Point
	}
	return nil
}

func(r *RollsBuilder) MakeSauce(sause string) error {
	if sause == "" {
		return errors.New(m.Wrong + m.Sauce)
	}
	r.Product.Ingredients += m.SauceInRolls + sause + m.Point
	return nil
}

func (p *Product) GetIngridients() (string, error) {
	if p.Ingredients == "" {
		return "", errors.New(m.Wrong + m.Ingridients)
	}
	return p.Ingredients, nil
}
