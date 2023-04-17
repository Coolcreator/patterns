package factory

import (
	"errors"
)

func GetPhone(phoneType string) (Phone, error) {
	switch phoneType {
	case "iphone":
		return &Iphone{phone{os: "IOS", manufacturer: "Apple"}}, nil
	case "sony":
		return &Sony{phone{os: "Android", manufacturer: "Sony"}}, nil
	default:
		return nil, errors.New("unknown phone type")
	}
}
