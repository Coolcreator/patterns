package director

import "patterns/internal/api/computer"

type computerBuilderMock struct {
	c computer.Computer
}

// SetProcessor устанавливает процессор
func (b *computerBuilderMock) SetProcessor() {
	b.c.Processor = "Test processor"
}

// SetGraphicsCard устанавливает графическую карту
func (b *computerBuilderMock) SetGraphicsCard() {
	b.c.GraphicsCard = "Test graphics card"
}

// SetRAM устанавливает оперативную память
func (b *computerBuilderMock) SetRAM() {
	b.c.RAM = "Test RAM"
}

// SetSSD устанавливает твердотельный накопитель
func (b *computerBuilderMock) SetSSD() {
	b.c.SSD = "Test SSD"
}

// SetOS устанавливает операционную систему
func (b *computerBuilderMock) SetOS() {
	b.c.OS = "Test OS"
}

// GetComputer возвращает экземпляр игрового компьютера
func (b *computerBuilderMock) GetComputer() computer.Computer {
	return b.c
}

// NewComputerBuilderMock возвращает экземпляр computerBuilderMock
func NewComputerBuilderMock() buildProcess {
	return &computerBuilderMock{}
}
