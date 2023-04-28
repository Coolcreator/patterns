package visitor

// Visitable представляет метод, с помощью которого его можно посетить
type Visitable interface {
	Accept(visitor VATVisitor)
}

// Product представляет собой товар
type product struct {
	name     string
	price    float64
	quantity int
}

// Accept принимает посетителя
func (p *product) Accept(visitor VATVisitor) {
	visitor.VisitProduct(p)
}

// NewProduct создает новый экземпляр продукта
func NewProduct(name string, price float64, quantity int) Visitable {
	return &product{
		name:     name,
		price:    price,
		quantity: quantity,
	}
}
