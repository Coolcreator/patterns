package director

import (
	"testing"
)

func TestDirector_SetBuilder(t *testing.T) {
	expected := NewComputerBuilderMock()
	d := &director{}
	d.SetBuilder(expected)
	got := d.builder
	if expected != got {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}

func TestDirector_BuildComputer(t *testing.T) {
	b := NewComputerBuilderMock()
	d := &director{builder: b}
	d.BuildComputer()
	c := b.GetComputer()
	if c.Processor != "Test processor" ||
		c.GraphicsCard != "Test graphics card" ||
		c.RAM != "Test RAM" ||
		c.SSD != "Test SSD" ||
		c.OS != "Test OS" {
		t.Errorf("build computer error: unexpected component")
	}
}
