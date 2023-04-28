package gaming

import (
	"testing"

	"patterns/internal/api/computer"
)

func TestGamingComputerBuilder_SetProcessor(t *testing.T) {
	gcb := &gamingComputerBuilder{}
	expected := "Intel Core i7 12th Gen 12700F"
	gcb.SetProcessor(expected)
	got := gcb.g.Processor
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestGamingComputerBuilder_SetGraphicsCard(t *testing.T) {
	gcb := &gamingComputerBuilder{}
	expected := "NVIDIA GeForce RTX 3070 8 GB"
	gcb.SetGraphicsCard(expected)
	got := gcb.g.GraphicsCard
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestGamingComputerBuilder_SetRAM(t *testing.T) {
	gcb := &gamingComputerBuilder{}
	expected := "16GB DDR4"
	gcb.SetRAM(expected)
	got := gcb.g.RAM
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestGamingComputerBuilder_SetSSD(t *testing.T) {
	gcb := &gamingComputerBuilder{}
	expected := "512 GB PCIe SSD"
	gcb.SetSSD(expected)
	got := gcb.g.SSD
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestGamingComputerBuilder_SetOS(t *testing.T) {
	gcb := &gamingComputerBuilder{}
	expected := "Windows 11 Home"
	gcb.SetOS(expected)
	got := gcb.g.OS
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestGamingComputerBuilder_GetComputer(t *testing.T) {
	expected := computer.Computer{}
	gcb := &gamingComputerBuilder{g: expected}
	got := gcb.GetComputer()
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}
