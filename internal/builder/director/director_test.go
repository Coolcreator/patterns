package director

import (
	"testing"

	"patterns/internal/builder/gaming"
	"patterns/internal/builder/working"
)

func TestDirector_SetBuilder(t *testing.T) {
	expected := gaming.NewGamingComputerBuilder()
	d := &director{}
	d.SetBuilder(expected)
	got := d.builder
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestDirector_BuildComputer(t *testing.T) {
	b := working.NewWorkingComputerBuilder()
	d := &director{builder: b}
	d.BuildComputer()
	c := b.GetComputer()
	if c.Processor == "" || c.GraphicsCard == "" || c.RAM == "" || c.SSD == "" || c.OS == "" {
		t.Errorf("build computer error: empty field")
	}
}
