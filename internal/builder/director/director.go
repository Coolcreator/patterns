package director

import "patterns/internal/builder"

// Director распоряжается строителем и задает ему порядок шагов строительства
type Director interface {
	SetBuilder(builder builder.BuildProcess)
	BuildComputer()
}

type director struct {
	builder builder.BuildProcess
}

// SetBuilder устанавливает строителя
func (d *director) SetBuilder(builder builder.BuildProcess) {
	d.builder = builder
}

// BuildComputer конструирует новый компьютер
func (d *director) BuildComputer() {
	d.builder.
		SetProcessor().
		SetGraphicsCard().
		SetRAM().
		SetSSD().
		SetOS()
}

// NewDirector создает новый экземпляр Director
func NewDirector(builder builder.BuildProcess) Director {
	return &director{
		builder: builder,
	}
}
