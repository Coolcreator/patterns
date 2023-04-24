package integration

import (
	"patterns/internal/api/computer"
	"testing"

	"patterns/internal/builder/director"
	"patterns/internal/builder/gaming"
	"patterns/internal/builder/working"
)

func TestBuilderPattern(t *testing.T) {
	expected := computer.Computer{
		Processor:    "Intel Core i7 12th Gen 12700F",
		GraphicsCard: "NVIDIA GeForce RTX 3070 8 GB",
		RAM:          "16GB DDR4",
		SSD:          "512 GB PCIe SSD",
		OS:           "Windows 11 Home",
	}
	gcb := gaming.NewGamingComputerBuilder()
	d := director.NewDirector(gcb)
	d.BuildComputer()
	got := gcb.GetComputer()
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
	expected = computer.Computer{
		Processor:    "Intel Core i7 8th Gen 8665U",
		GraphicsCard: "Intel UHD Graphics 620",
		RAM:          "8GB DDR4",
		SSD:          "512 GB PCIe SSD",
		OS:           "Windows 11 Pro",
	}
	wcb := working.NewWorkingComputerBuilder()
	d.SetBuilder(wcb)
	d.BuildComputer()
	got = wcb.GetComputer()
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}
