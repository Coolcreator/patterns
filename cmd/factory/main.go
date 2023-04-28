package main

import (
	"fmt"

	"patterns/internal/factory"
)

func main() {
	f := factory.NewLaptopFactory()
	l := f.CreateLaptop("lenovo", "thinkPad")
	fmt.Println(l.DisplayDescription())
	d := f.CreateLaptop("dell", "latitude")
	fmt.Println(d.DisplayDescription())
}
