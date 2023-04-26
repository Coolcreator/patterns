package working

import (
	"patterns/internal/api/computer"
)

// WorkingComputerBuilder представляет собой билдер рабочего компьютера
type WorkingComputerBuilder interface {
	SetProcessor()
	SetGraphicsCard()
	SetRAM()
	SetSSD()
	SetOS()
	GetComputer() computer.Computer
}

type workingComputerBuilder struct {
	w computer.Computer
}

// SetProcessor устанавливает процессор
func (b *workingComputerBuilder) SetProcessor() {
	b.w.Processor = "Intel Core i7 8th Gen 8665U"
}

// SetGraphicsCard устанавливает графическую карту
func (b *workingComputerBuilder) SetGraphicsCard() {
	b.w.GraphicsCard = "Intel UHD Graphics 620"
}

// SetRAM устанавливает оперативную память
func (b *workingComputerBuilder) SetRAM() {
	b.w.RAM = "8GB DDR4"
}

// SetSSD устанавливает твердотельный накопитель
func (b *workingComputerBuilder) SetSSD() {
	b.w.SSD = "512 GB PCIe SSD"
}

// SetOS устанавливает операционную систему
func (b *workingComputerBuilder) SetOS() {
	b.w.OS = "Windows 11 Pro"
}

// GetComputer возвращает экземпляр рабочего компьютера
func (b *workingComputerBuilder) GetComputer() computer.Computer {
	return b.w
}

// NewWorkingComputerBuilder возвращает экземпляр workingComputerBuilder
func NewWorkingComputerBuilder() WorkingComputerBuilder {
	return &workingComputerBuilder{}
}
