package working

import (
	"patterns/internal/api/computer"
)

// WorkingComputerBuilder представляет собой билдер рабочего компьютера
type WorkingComputerBuilder interface {
	SetProcessor(processor string)
	SetGraphicsCard(graphicsCard string)
	SetRAM(ram string)
	SetSSD(ssd string)
	SetOS(os string)
	GetComputer() computer.Computer
}

type workingComputerBuilder struct {
	w computer.Computer
}

// SetProcessor устанавливает процессор
func (b *workingComputerBuilder) SetProcessor(processor string) {
	b.w.Processor = processor
}

// SetGraphicsCard устанавливает графическую карту
func (b *workingComputerBuilder) SetGraphicsCard(graphicsCard string) {
	b.w.GraphicsCard = graphicsCard
}

// SetRAM устанавливает оперативную память
func (b *workingComputerBuilder) SetRAM(ram string) {
	b.w.RAM = ram
}

// SetSSD устанавливает твердотельный накопитель
func (b *workingComputerBuilder) SetSSD(SSD string) {
	b.w.SSD = SSD
}

// SetOS устанавливает операционную систему
func (b *workingComputerBuilder) SetOS(OS string) {
	b.w.OS = OS
}

// GetComputer возвращает экземпляр рабочего компьютера
func (b *workingComputerBuilder) GetComputer() computer.Computer {
	return b.w
}

// NewWorkingComputerBuilder возвращает экземпляр workingComputerBuilder
func NewWorkingComputerBuilder() WorkingComputerBuilder {
	return &workingComputerBuilder{}
}
