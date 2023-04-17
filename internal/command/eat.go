package command

import "fmt"

type Eat struct {
	Food string
}

func (e *Eat) execute() {
	fmt.Printf("ate %s\n", e.Food)
}
