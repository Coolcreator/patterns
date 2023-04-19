package asylum

import (
	"fmt"

	cor "patterns/internal/chain-of-responsibility"
	"patterns/internal/chain-of-responsibility/api/models"
)

type asylum struct {
	next cor.Step
}

func (a *asylum) Execute(p *models.Person) {
	fmt.Printf("[Asylum] person %s is in mental institution\n", p.Name)
}

func (a *asylum) SetNext(next cor.Step) {
	a.next = next
}

func NewAsylum() cor.Step {
	return &asylum{}
}
