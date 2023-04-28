package work

import (
	"fmt"

	"patterns/internal/api/person"
)

type step interface {
	Execute(p *person.Person)
}

// Work представляет собой обработчик работы
type Work interface {
	step
	SetNext(next step)
}

type work struct {
	next step
}

// Execute запускает обработчик работы
func (w *work) Execute(p *person.Person) {
	fmt.Printf("[Work] person %s is working\n", p.Name)

	w.next.Execute(p)
}

// SetNext устанавливает следующий обработчик
func (w *work) SetNext(next step) {
	w.next = next
}

// NewWork возвращает экземпляр обработчика работы
func NewWork() Work {
	return &work{}
}
