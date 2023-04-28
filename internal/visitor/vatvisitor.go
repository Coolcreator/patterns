package visitor

import "fmt"

// VATVisitor представляет собой посетителя для рассчета НДС
type VATVisitor interface {
	VisitProduct(product *product)
}

type vatVisitor struct {
	vatRate float64
}

func (v *vatVisitor) VisitProduct(product *product) {
	vat := (product.price * float64(product.quantity)) * v.vatRate
	fmt.Printf("Товар: %s, НДС: %0.2f\n", product.name, vat)
}

// NewVATVisitor создает новый экземпляр VATVisitor
func NewVATVisitor(vatRate float64) VATVisitor {
	return &vatVisitor{vatRate: vatRate}
}
