package builder

type Director struct {
	builder ComputerBuilder
}

func (d *Director) SetBuilder(builder ComputerBuilder) {
	d.builder = builder
}

func (d *Director) BuildComputer() Computer {
	d.builder.SetProcessor()
	d.builder.SetGraphicsCard()
	d.builder.SetRAM()
	d.builder.SetSSD()
	d.builder.SetOS()

	return d.builder.GetComputer()
}
