package university

import (
	"fmt"

	cor "patterns/internal/chain-of-responsibility"
	"patterns/internal/chain-of-responsibility/api/models"
)

type university struct {
	next cor.Step
}

func (u *university) Execute(p *models.Person) {
	fmt.Printf("[University] person %s is learning in university\n", p.Name)

	u.next.Execute(p)
}

func (u *university) SetNext(next cor.Step) {
	u.next = next
}

func NewUniversity() cor.Step {
	return &university{}
}
