package factory

type Iphone struct {
	phone
}

func (i *Iphone) GetOS() string {
	return i.os
}

func (i *Iphone) GetManufacturer() string {
	return i.manufacturer
}
