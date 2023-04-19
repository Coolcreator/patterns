package main

import (
	"log"

	"patterns/internal/factory/factory"
)

func main() {
	log.SetFlags(0)

	iphone, err := factory.GetPhone("iphone")
	if err != nil {
		log.Fatalf("get phone error: %v\n", err)
	}
	log.Println("OS:", iphone.GetOS(), "Manufacturer:", iphone.GetManufacturer())

	sony, err := factory.GetPhone("sony")
	if err != nil {
		log.Fatalf("get phone error: %v\n", err)
	}
	log.Println("OS:", sony.GetOS(), "Manufacturer:", sony.GetManufacturer())

	blackberry, err := factory.GetPhone("blackberry")
	if err != nil {
		log.Fatalf("get phone error: %v\n", err)
	}
	log.Println("OS:", blackberry.GetOS(), "Manufacturer:", blackberry.GetManufacturer())
}
