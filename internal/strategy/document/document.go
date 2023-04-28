package document

import "fmt"

// Document представляет собой документ
type Document interface {
	Print()
}

type document struct {
	format string
}

// Print печатает документ
func (d *document) Print() {
	fmt.Println("printing document with format", d.format)
}

// NewDocument создает новый документ
func NewDocument(format string) Document {
	return &document{format: format}
}
