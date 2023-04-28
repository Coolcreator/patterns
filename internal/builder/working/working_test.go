package working

import (
	"testing"

	"patterns/internal/api/computer"
)

func TestWorkingComputerBuilder_SetProcessor(t *testing.T) {
	wcb := &workingComputerBuilder{}
	expected := "Intel Core i7 8th Gen 8665U"
	wcb.SetProcessor(expected)
	got := wcb.w.Processor
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestWorkingComputerBuilder_SetGraphicsCard(t *testing.T) {
	wcb := &workingComputerBuilder{}
	expected := "Intel UHD Graphics 620"
	wcb.SetGraphicsCard(expected)
	got := wcb.w.GraphicsCard
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestWorkingComputerBuilder_SetRAM(t *testing.T) {
	wcb := &workingComputerBuilder{}
	expected := "8GB DDR4"
	wcb.SetRAM(expected)
	got := wcb.w.RAM
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestWorkingComputerBuilder_SetSSD(t *testing.T) {
	wcb := &workingComputerBuilder{}
	expected := "512 GB PCIe SSD"
	wcb.SetSSD(expected)
	got := wcb.w.SSD
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestWorkingComputerBuilder_SetOS(t *testing.T) {
	wcb := &workingComputerBuilder{}
	expected := "Windows 11 Pro"
	wcb.SetOS(expected)
	got := wcb.w.OS
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestWorkingComputerBuilder_GetComputer(t *testing.T) {
	expected := computer.Computer{}
	wcb := &workingComputerBuilder{w: expected}
	got := wcb.GetComputer()
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}
