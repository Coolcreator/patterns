package visitor

import "fmt"

type CalculatePrice struct {
	Additional float32
}

func (p CalculatePrice) VisitBread(b Bread) {
	fmt.Printf("bread price now is %0.2f\n", b.price+p.Additional)
}

func (p CalculatePrice) VisitCheese(c Cheese) {
	fmt.Printf("cheese price now is %0.2f\n", c.price+p.Additional)
}
