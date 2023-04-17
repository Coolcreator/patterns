package visitor

import "fmt"

type Cheese struct {
	name  string
	price float32
}

func NewCheese(name string, price float32) *Cheese {
	return &Cheese{name: name, price: price}
}

func (c Cheese) Name() {
	fmt.Println(c.name)
}

func (c Cheese) Accept(v Visitor) {
	v.VisitCheese(c)
}
