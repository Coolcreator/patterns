package main

import (
	"patterns/internal/api/person"
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
	w.SetNext(a)
	u.SetNext(w)
	s.SetNext(u)
	p := &person.Person{
		Name:   "Aleksey",
		Course: 1,
		Group:  118,
	}
	s.Execute(p)
}
