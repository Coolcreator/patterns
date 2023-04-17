package main

import (
	cor "patterns/internal/chain-of-responsibility"
)

func main() {
	asylum := &cor.Asylum{}

	work := &cor.Work{}
	work.SetNext(asylum)

	university := &cor.University{}
	university.SetNext(work)

	school := &cor.School{}
	school.SetNext(university)

	p := &cor.Person{Name: "Alexey"}

	school.Execute(p)
}
