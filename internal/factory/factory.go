package factory

// Laptop представляет собой ноутбук
type Laptop interface {
	DisplayDescription() string
}

// LaptopFactory представляет собой фабрику ноутбуков
type LaptopFactory interface {
	CreateLaptop(brand string, name string) Laptop
}

type laptopFactory struct{}

// CreateLaptop создает ноутбук в зависимости от производителя
func (f laptopFactory) CreateLaptop(brand string, name string) Laptop {
	switch brand {
	case "lenovo":
		return &lenovo{Name: name}
	case "dell":
		return &dell{Name: name}
	default:
		return nil
	}
}

// NewLaptopFactory создает новый экземпляр LaptopFactory
func NewLaptopFactory() LaptopFactory {
	return &laptopFactory{}
}
