package visitor

import "fmt"

type Bread struct {
	name  string
	price float32
}

func NewBread(name string, price float32) *Bread {
	return &Bread{name: name, price: price}
}

func (b Bread) Name() {
	fmt.Println(b.name)
}

func (b Bread) Accept(v Visitor) {
	v.VisitBread(b)
}
