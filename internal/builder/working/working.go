package working

import (
	"patterns/internal/api/computer"
	"patterns/internal/builder"
)

type workingComputerBuilder struct {
	w computer.Computer
}

// SetProcessor устанавливает процессор
func (b *workingComputerBuilder) SetProcessor() builder.BuildProcess {
	b.w.Processor = "Intel Core i7 8th Gen 8665U"
	return b
}

// SetGraphicsCard устанавливает графическую карту
func (b *workingComputerBuilder) SetGraphicsCard() builder.BuildProcess {
	b.w.GraphicsCard = "Intel UHD Graphics 620"
	return b
}

// SetRAM устанавливает оперативную память
func (b *workingComputerBuilder) SetRAM() builder.BuildProcess {
	b.w.RAM = "8GB DDR4"
	return b
}

// SetSSD устанавливает твердотельный накопитель
func (b *workingComputerBuilder) SetSSD() builder.BuildProcess {
	b.w.SSD = "512 GB PCIe SSD"
	return b
}

// SetOS устанавливает операционную систему
func (b *workingComputerBuilder) SetOS() builder.BuildProcess {
	b.w.OS = "Windows 11 Pro"
	return b
}

// GetComputer возвращает экземпляр рабочего компьютера
func (b *workingComputerBuilder) GetComputer() computer.Computer {
	return b.w
}

// NewWorkingComputerBuilder возвращает экземпляр workingComputerBuilder
func NewWorkingComputerBuilder() *workingComputerBuilder {
	return &workingComputerBuilder{}
}
