package factory

type Sony struct {
	phone
}

func (s *Sony) GetOS() string {
	return s.os
}

func (s *Sony) GetManufacturer() string {
	return s.manufacturer
}
