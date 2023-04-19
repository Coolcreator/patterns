package factory

import (
	"errors"

	"patterns/internal/factory/iphone"
	"patterns/internal/factory/sony"
)

type Phone interface {
	GetOS() string
	GetManufacturer() string
}

func GetPhone(phoneType string) (Phone, error) {
	switch phoneType {
	case "iphone":
		return iphone.NewPhone(), nil
	case "sony":
		return sony.NewPhone(), nil
	default:
		return nil, errors.New("unknown phone type")
	}
}
