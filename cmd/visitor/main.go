package main

import "patterns/internal/visitor"

func main() {
	bread := visitor.NewBread("white bread with oregano", 12.0)
	cheese := visitor.NewCheese("dutch cheese", 250.0)

	bread.Name()
	bread.Accept(&visitor.CalculatePrice{Additional: 3.0})

	cheese.Name()
	cheese.Accept(&visitor.CalculatePrice{Additional: 50.0})
}
