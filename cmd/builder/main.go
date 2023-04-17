package main

import (
	"fmt"
	"patterns/internal/builder"
)

func main() {
	gamingComputer := &builder.GamingComputerBuilder{}
	workingComputer := &builder.WorkingComputerBuilder{}
	director := &builder.Director{}
	director.SetBuilder(gamingComputer)
	computer := director.BuildComputer()
	fmt.Println(computer)
	director.SetBuilder(workingComputer)
	computer = director.BuildComputer()
	fmt.Println(computer)
}
