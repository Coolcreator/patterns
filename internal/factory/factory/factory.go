package factory

import (
	"errors"
)

// Phone представляет собой тип телефона
type Phone interface {
	GetOS() string
	GetManufacturer() string
}

// Factory представляет собой фабрику телефонов
type Factory interface {
	GetPhone(phoneType string) (Phone, error)
}

type factory struct {
	phones map[string]Phone
}

// GetPhone конструирует телефон
func (f *factory) GetPhone(phoneType string) (Phone, error) {
	p, ok := f.phones[phoneType]
	if !ok {
		return nil, errors.New("unknown phone type")
	}
	return p, nil
}

// NewFactory создает новую фабрику телефонов
func NewFactory(phones map[string]Phone) Factory {
	return &factory{phones: phones}
}
