package school

import (
	"fmt"

	"patterns/internal/api/person"
)

type step interface {
	Execute(p *person.Person)
}

// School представляет собой обработчик школы
type School interface {
	step
	SetNext(next step)
}

type school struct {
	next step
}

// Execute запускает обработчик школы
func (s *school) Execute(p *person.Person) {
	fmt.Printf("[School] person %s is learning in school\n", p.Name)

	s.next.Execute(p)
}

// SetNext устанавливает следующий обработчик
func (s *school) SetNext(next step) {
	s.next = next
}

// NewSchool возвращает экземпляр обработчика школы
func NewSchool() School {
	return &school{}
}
