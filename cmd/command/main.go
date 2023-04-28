package main

import (
	"fmt"

	"patterns/internal/command/calculator"
	"patterns/internal/command/command"
	"patterns/internal/command/director"
)

func main() {
	c := calculator.NewCalculator()
	director := director.NewDirector()
	director.ExecuteCommand(command.NewAddCommand(c, 10))
	fmt.Println(c.GetValue())
	director.ExecuteCommand(command.NewSubtractCommand(c, 5))
	fmt.Println(c.GetValue())
	director.UndoLastCommand()
	fmt.Println(c.GetValue())
}
