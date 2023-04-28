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
	wpc := d.BuildComputer(
		"Intel Core i7 8th Gen 8665U",
		"Intel UHD Graphics 620",
		"8GB DDR4",
		"512 GB PCIe SSD",
		"Windows 11 Pro",
	)
	fmt.Println("Working computer setup:", wpc)
	gcb := gaming.NewGamingComputerBuilder()
	d.SetBuilder(gcb)
	gpc := d.BuildComputer(
		"Intel Core i7 12th Gen 12700F",
		"NVIDIA GeForce RTX 3070 8 GB",
		"16GB DDR4",
		"512 GB PCIe SSD",
		"Windows 11 Home",
	)
	fmt.Println("Gaming computer setup:", gpc)
}
