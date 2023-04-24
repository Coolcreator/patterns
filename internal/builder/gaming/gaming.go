package gaming

import (
	"patterns/internal/api/computer"
	"patterns/internal/builder"
)

type gamingComputerBuilder struct {
	g computer.Computer
}

// SetProcessor устанавливает процессор
func (b *gamingComputerBuilder) SetProcessor() builder.BuildProcess {
	b.g.Processor = "Intel Core i7 12th Gen 12700F"
	return b
}

// SetGraphicsCard устанавливает графическую карту
func (b *gamingComputerBuilder) SetGraphicsCard() builder.BuildProcess {
	b.g.GraphicsCard = "NVIDIA GeForce RTX 3070 8 GB"
	return b
}

// SetRAM устанавливает оперативную память
func (b *gamingComputerBuilder) SetRAM() builder.BuildProcess {
	b.g.RAM = "16GB DDR4"
	return b
}

// SetSSD устанавливает твердотельный накопитель
func (b *gamingComputerBuilder) SetSSD() builder.BuildProcess {
	b.g.SSD = "512 GB PCIe SSD"
	return b
}

// SetOS устанавливает операционную систему
func (b *gamingComputerBuilder) SetOS() builder.BuildProcess {
	b.g.OS = "Windows 11 Home"
	return b
}

// GetComputer возвращает экземпляр игрового компьютера
func (b *gamingComputerBuilder) GetComputer() computer.Computer {
	return b.g
}

// NewGamingComputerBuilder возвращает экземпляр gamingComputerBuilder
func NewGamingComputerBuilder() *gamingComputerBuilder {
	return &gamingComputerBuilder{}
}
