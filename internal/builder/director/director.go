package director

import (
	"patterns/internal/builder/api/models"
)

type computerBuilder interface {
	SetProcessor()
	SetGraphicsCard()
	SetRAM()
	SetMemory()
	SetOS()
	GetComputer() models.Computer
}

type Director interface {
	SetBuilder(builder computerBuilder)
	BuildComputer() models.Computer
}

type director struct {
	builder computerBuilder
}

func (d *director) SetBuilder(builder computerBuilder) {
	d.builder = builder
}

func (d *director) BuildComputer() models.Computer {
	d.builder.SetProcessor()
	d.builder.SetGraphicsCard()
	d.builder.SetRAM()
	d.builder.SetMemory()
	d.builder.SetOS()

	return d.builder.GetComputer()
}

func NewDirector(builder computerBuilder) Director {
	return &director{
		builder: builder,
	}
}
