package director

import (
	"patterns/internal/api/computer"
)

type buildProcess interface {
	SetProcessor()
	SetGraphicsCard()
	SetRAM()
	SetSSD()
	SetOS()
	GetComputer() computer.Computer
}

// Director распоряжается строителем и задает ему порядок шагов строительства
type Director interface {
	SetBuilder(builder buildProcess)
	BuildComputer()
}

type director struct {
	builder buildProcess
}

// SetBuilder устанавливает строителя
func (d *director) SetBuilder(builder buildProcess) {
	d.builder = builder
}

// BuildComputer конструирует новый компьютер
func (d *director) BuildComputer() {
	d.builder.SetProcessor()
	d.builder.SetGraphicsCard()
	d.builder.SetRAM()
	d.builder.SetSSD()
	d.builder.SetOS()
}

// NewDirector создает новый экземпляр Director
func NewDirector(builder buildProcess) Director {
	return &director{
		builder: builder,
	}
}
