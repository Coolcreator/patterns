package eat

import "fmt"

// Eat представляет собой команду "кушать"
type Eat interface {
	Execute()
}

type eat struct {
	food string
}

// Execute выполняет команду "кушать"
func (e *eat) Execute() {
	fmt.Printf("ate %s\n", e.food)
}

// NewEat создает экземпляр команды "кушать"
func NewEat(food string) Eat {
	return &eat{food: food}
}
