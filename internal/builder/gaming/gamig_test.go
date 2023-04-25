package gaming

import (
	"testing"

	"patterns/internal/api/computer"
)

func TestGamingComputerBuilder_SetProcessor(t *testing.T) {
	gcb := &gamingComputerBuilder{}
	gcb.SetProcessor()
	expected := "Intel Core i7 12th Gen 12700F"
	got := gcb.g.Processor
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestGamingComputerBuilder_SetGraphicsCard(t *testing.T) {
	gcb := &gamingComputerBuilder{}
	gcb.SetGraphicsCard()
	expected := "NVIDIA GeForce RTX 3070 8 GB"
	got := gcb.g.GraphicsCard
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestGamingComputerBuilder_SetRAM(t *testing.T) {
	gcb := &gamingComputerBuilder{}
	gcb.SetRAM()
	expected := "16GB DDR4"
	got := gcb.g.RAM
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestGamingComputerBuilder_SetSSD(t *testing.T) {
	gcb := &gamingComputerBuilder{}
	gcb.SetSSD()
	expected := "512 GB PCIe SSD"
	got := gcb.g.SSD
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestGamingComputerBuilder_SetOS(t *testing.T) {
	gcb := &gamingComputerBuilder{}
	gcb.SetOS()
	expected := "Windows 11 Home"
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
