package sony

import (
	"patterns/internal/api/phone"
)

// Sony представляет собой смартфон Sony Xperia
type Sony interface {
	GetOS() string
	GetManufacturer() string
}

type sony struct {
	p phone.Phone
}

// GetOS возвращает операционную систему смартфона
func (s *sony) GetOS() string {
	return s.p.OS
}

// GetManufacturer возвращает производителя смартфона
func (s *sony) GetManufacturer() string {
	return s.p.Manufacturer
}

// NewSony создает новый экземпляр смартфона Sony
func NewSony() Sony {
	return &sony{
		p: phone.Phone{
			OS:           "Android",
			Manufacturer: "Sony",
		},
	}
}
