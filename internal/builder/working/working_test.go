package working

import (
	"testing"

	"patterns/internal/api/computer"
)

func TestWorkingComputerBuilder_SetProcessor(t *testing.T) {
	wcb := &workingComputerBuilder{}
	wcb.SetProcessor()
	expected := "Intel Core i7 8th Gen 8665U"
	got := wcb.w.Processor
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestWorkingComputerBuilder_SetGraphicsCard(t *testing.T) {
	wcb := &workingComputerBuilder{}
	wcb.SetGraphicsCard()
	expected := "Intel UHD Graphics 620"
	got := wcb.w.GraphicsCard
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestWorkingComputerBuilder_SetRAM(t *testing.T) {
	wcb := &workingComputerBuilder{}
	wcb.SetRAM()
	expected := "8GB DDR4"
	got := wcb.w.RAM
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestWorkingComputerBuilder_SetSSD(t *testing.T) {
	wcb := &workingComputerBuilder{}
	wcb.SetSSD()
	expected := "512 GB PCIe SSD"
	got := wcb.w.SSD
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestWorkingComputerBuilder_SetOS(t *testing.T) {
	wcb := &workingComputerBuilder{}
	wcb.SetOS()
	expected := "Windows 11 Pro"
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
