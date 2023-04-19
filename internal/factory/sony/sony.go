package sony

import "patterns/internal/factory/api/models"

type phone struct {
	models.Phone
}

func (s *phone) GetOS() string {
	return s.OS
}

func (s *phone) GetManufacturer() string {
	return s.Manufacturer
}

func NewPhone() *phone {
	return &phone{
		Phone: models.Phone{
			OS:           "Android",
			Manufacturer: "Sony",
		},
	}
}
