package gaming

import (
	"patterns/internal/api/computer"
)

type gamingComputerBuilder struct {
	g computer.Computer
}

// SetProcessor устанавливает процессор
func (b *gamingComputerBuilder) SetProcessor() {
	b.g.Processor = "Intel Core i7 12th Gen 12700F"
}

// SetGraphicsCard устанавливает графическую карту
func (b *gamingComputerBuilder) SetGraphicsCard() {
	b.g.GraphicsCard = "NVIDIA GeForce RTX 3070 8 GB"
}

// SetRAM устанавливает оперативную память
func (b *gamingComputerBuilder) SetRAM() {
	b.g.RAM = "16GB DDR4"
}

// SetSSD устанавливает твердотельный накопитель
func (b *gamingComputerBuilder) SetSSD() {
	b.g.SSD = "512 GB PCIe SSD"
}

// SetOS устанавливает операционную систему
func (b *gamingComputerBuilder) SetOS() {
	b.g.OS = "Windows 11 Home"
}

// GetComputer возвращает экземпляр игрового компьютера
func (b *gamingComputerBuilder) GetComputer() computer.Computer {
	return b.g
}

// NewGamingComputerBuilder возвращает экземпляр gamingComputerBuilder
func NewGamingComputerBuilder() *gamingComputerBuilder {
	return &gamingComputerBuilder{}
}
