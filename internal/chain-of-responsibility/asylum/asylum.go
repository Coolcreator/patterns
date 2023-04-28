package asylum

import (
	"fmt"

	"patterns/internal/api/person"
)

type executor interface {
	Execute(p *person.Person)
}

// Asylum представляет собой обработчик психиатрической больницы
type Asylum interface {
	executor
	SetNext(next executor)
}

type asylum struct {
	next executor
}

// Execute запускает обработчик психиатрической больницы
func (*asylum) Execute(p *person.Person) {
	fmt.Printf("[Asylum] person %s is in mental institution\n", p.Name)
}

// SetNext устанавливает следующий обработчик
func (a *asylum) SetNext(next executor) {
	a.next = next
}

// NewAsylum возвращает обработчик психиатрической больницы
func NewAsylum() Asylum {
	return &asylum{}
}
