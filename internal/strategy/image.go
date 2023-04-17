package strategy

import "fmt"

type Image struct {
	format string
}

func NewImage(format string) *Image {
	return &Image{format: format}
}

func (i *Image) print() {
	fmt.Println("printing image: format", i.format)
}

func (i *Image) fmt() {
	fmt.Println("formatting image...")
}
