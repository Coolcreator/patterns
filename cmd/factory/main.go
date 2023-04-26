package main

import (
	"log"

	"patterns/internal/factory/factory"
	"patterns/internal/factory/iphone"
	"patterns/internal/factory/sony"
)

func main() {
	log.SetFlags(0)
	f := factory.NewFactory(map[string]factory.Phone{
		"sony":   sony.NewSony(),
		"iphone": iphone.NewIPhone(),
	})
	p, err := f.GetPhone("sony")
	if err != nil {
		log.Fatalf("get phone error: %v\n", err)
	}
	log.Println("OS:", p.GetOS(), "Manufacturer:", p.GetManufacturer())
	p, err = f.GetPhone("iphone")
	if err != nil {
		log.Fatalf("get phone error: %v\n", err)
	}
	log.Println("OS:", p.GetOS(), "Manufacturer:", p.GetManufacturer())
}
