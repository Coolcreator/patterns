package image

import "fmt"

// Image представляет собой изображение
type Image interface {
	Print()
}

type image struct {
	format string
}

// Print печатает изображение
func (i *image) Print() {
	fmt.Println("printing image with format", i.format)
}

// NewImage создает новое изображение
func NewImage(format string) Image {
	return &image{format: format}
}
