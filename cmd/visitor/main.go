package main

import "patterns/internal/visitor"

func main() {
	product := visitor.NewProduct("Шоколад", 100, 2)
	vatVisitor := visitor.NewVATVisitor(0.20)
	product.Accept(vatVisitor)
}
