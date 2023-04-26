package iphone

import (
	"patterns/internal/api/phone"
)

// Iphone представляет собой смартфон Apple Iphone
type Iphone interface {
	GetOS() string
	GetManufacturer() string
}

type iphone struct {
	p phone.Phone
}

// GetOS возвращает операционную систему смартфона
func (i *iphone) GetOS() string {
	return i.p.OS
}

// GetManufacturer возвращает производителя смартфона
func (i *iphone) GetManufacturer() string {
	return i.p.Manufacturer
}

// NewIPhone создает новый экземпляр смартфона Apple Iphone
func NewIPhone() Iphone {
	return &iphone{
		p: phone.Phone{
			OS:           "IOS",
			Manufacturer: "Apple",
		},
	}
}
