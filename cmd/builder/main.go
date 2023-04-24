package main

import (
	"fmt"

	"patterns/internal/builder/director"
	"patterns/internal/builder/gaming"
	"patterns/internal/builder/working"
)

func main() {
	wcb := working.NewWorkingComputerBuilder()
	d := director.NewDirector(wcb)
	d.BuildComputer()
	wpc := wcb.GetComputer()
	fmt.Println("Working computer setup:", wpc)
	gcb := gaming.NewGamingComputerBuilder()
	d.SetBuilder(gcb)
	d.BuildComputer()
	gpc := gcb.GetComputer()
	fmt.Println("Gaming computer setup:", gpc)
}
