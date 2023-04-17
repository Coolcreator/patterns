package visitor

type Food interface {
	Name()
	Accept(Visitor)
}
