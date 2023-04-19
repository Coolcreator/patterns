package main

import (
	"patterns/internal/chain-of-responsibility/api/models"
	"patterns/internal/chain-of-responsibility/asylum"
	"patterns/internal/chain-of-responsibility/school"
	"patterns/internal/chain-of-responsibility/university"
	"patterns/internal/chain-of-responsibility/work"
)

func main() {
	a := asylum.NewAsylum()
	w := work.NewWork()
	u := university.NewUniversity()
	s := school.NewSchool()

	s.SetNext(u)
	u.SetNext(w)
	w.SetNext(a)

	p := &models.Person{
		Name: "Ivar",
	}

	s.Execute(p)
}
