package work

import (
	"fmt"

	cor "patterns/internal/chain-of-responsibility"
	"patterns/internal/chain-of-responsibility/api/models"
)

type work struct {
	next cor.Step
}

func (w *work) Execute(p *models.Person) {
	fmt.Printf("[Work] person %s is working\n", p.Name)

	w.next.Execute(p)
}

func (w *work) SetNext(next cor.Step) {
	w.next = next
}

func NewWork() cor.Step {
	return &work{}
}
