package factory

type phone struct {
	os           string
	manufacturer string
}

type Phone interface {
	GetOS() string
	GetManufacturer() string
}
