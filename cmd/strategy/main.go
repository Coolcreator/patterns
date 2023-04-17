package main

import "patterns/internal/strategy"

func main() {
	img := strategy.NewImage("png")
	doc := strategy.NewDocument("pdf")

	strategy.Output(img)
	strategy.Output(doc)
}
