package main

import (
	"patterns/internal/command/code"
	"patterns/internal/command/developer"
	"patterns/internal/command/eat"
	"patterns/internal/command/sleep"
)

func main() {
	d := developer.NewDeveloper()
	e := eat.NewEat("pie")
	d.SetCommand(e).Do()
	s := sleep.NewSleep(60)
	d.SetCommand(s).Do()
	c := code.NewCode("recipeService")
	d.SetCommand(c).Do()
}
