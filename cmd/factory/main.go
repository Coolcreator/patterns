package main

import (
	"fmt"
	"patterns/internal/factory"
)

func main() {
	phone, _ := factory.GetPhone("iphone")
	fmt.Println(phone.GetOS(), phone.GetManufacturer())
	phone, _ = factory.GetPhone("sony")
	fmt.Println(phone.GetOS(), phone.GetManufacturer())
	phone, err := factory.GetPhone("?")
	fmt.Println(err)
}
