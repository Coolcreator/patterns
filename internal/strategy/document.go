package strategy

import "fmt"

type document struct {
	format string
}

func NewDocument(format string) *document {
	return &document{}
}

func (d *document) print() {
	fmt.Println("printing document: format", d.format)
}

func (d *document) fmt() {
	fmt.Println("formatting document...")
}
