package main

import (
	"fmt"

	"patterns/internal/builder/director"
	"patterns/internal/builder/game"
	"patterns/internal/builder/work"
)

func main() {
	wcb := work.NewWorkingComputerBuilder()
	d := director.NewDirector(wcb)
	pc := d.BuildComputer()
	fmt.Println("Working computer setup:", pc)

	gcb := game.NewGamingComputerBuilder()
	d.SetBuilder(gcb)
	pc = d.BuildComputer()
	fmt.Println("Gaming computer setup:", pc)
}
