package school

import (
	"fmt"

	cor "patterns/internal/chain-of-responsibility"
	"patterns/internal/chain-of-responsibility/api/models"
)

type school struct {
	next cor.Step
}

func (s *school) Execute(p *models.Person) {
	fmt.Printf("[School] person %s is learning in school\n", p.Name)

	s.next.Execute(p)
}

func (s *school) SetNext(next cor.Step) {
	s.next = next
}

func NewSchool() cor.Step {
	return &school{}
}
