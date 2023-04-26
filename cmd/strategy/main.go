package main

import (
	"patterns/internal/strategy/document"
	"patterns/internal/strategy/image"
	"patterns/internal/strategy/printer"
)

func main() {
	img := image.NewImage("png")
	printer.Print(img)
	doc := document.NewDocument("pdf")
	printer.Print(doc)
}
