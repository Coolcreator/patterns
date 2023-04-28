package director

import (
	"patterns/internal/api/computer"
)

type buildProcess interface {
	SetProcessor(processor string)
	SetGraphicsCard(graphicsCard string)
	SetRAM(ram string)
	SetSSD(ssd string)
	SetOS(os string)
	GetComputer() computer.Computer
}

// Director распоряжается строителем и задает ему порядок шагов строительства
type Director interface {
	SetBuilder(builder buildProcess)
	BuildComputer(processor, graphicsCard, ram, ssd, os string) computer.Computer
}

type director struct {
	builder buildProcess
}

// SetBuilder устанавливает строителя
func (d *director) SetBuilder(builder buildProcess) {
	d.builder = builder
}

// BuildComputer конструирует новый компьютер
func (d *director) BuildComputer(
	processor string,
	graphicsCard string,
	ram string,
	ssd string,
	os string,
) computer.Computer {
	d.builder.SetProcessor(processor)
	d.builder.SetGraphicsCard(graphicsCard)
	d.builder.SetRAM(ram)
	d.builder.SetSSD(ssd)
	d.builder.SetOS(os)
	return d.builder.GetComputer()
}

// NewDirector создает новый экземпляр Director
func NewDirector(builder buildProcess) Director {
	return &director{
		builder: builder,
	}
}
