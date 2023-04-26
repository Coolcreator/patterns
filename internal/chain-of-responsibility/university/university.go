package university

import (
	"fmt"

	"patterns/internal/api/person"
)

type step interface {
	Execute(p *person.Person)
}

// University представляет собой обработчик университета
type University interface {
	step
	SetNext(next step)
}

type university struct {
	next step
}

// Execute запускает обработчик университета
func (u *university) Execute(p *person.Person) {
	fmt.Printf("[University] person %s is learning in university\n", p.Name)

	u.next.Execute(p)
}

// SetNext устанавливает следующий обработчик
func (u *university) SetNext(next step) {
	u.next = next
}

// NewUniversity возвращает экземпляр обработчика университета
func NewUniversity() University {
	return &university{}
}
