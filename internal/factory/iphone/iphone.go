package iphone

import "patterns/internal/factory/api/models"

type phone struct {
	models.Phone
}

func (p *phone) GetOS() string {
	return p.OS
}

func (p *phone) GetManufacturer() string {
	return p.Manufacturer
}

func NewPhone() *phone {
	return &phone{
		Phone: models.Phone{
			OS:           "IOS",
			Manufacturer: "Apple",
		},
	}
}
