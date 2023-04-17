package visitor

type Visitor interface {
	VisitBread(b Bread)
	VisitCheese(c Cheese)
}
