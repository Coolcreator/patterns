package gaming

import (
	"patterns/internal/api/computer"
)

// GamingComputerBuilder представляет собой билдер игрового компьютера
type GamingComputerBuilder interface {
	SetProcessor(processor string)
	SetGraphicsCard(graphicsCard string)
	SetRAM(ram string)
	SetSSD(ssd string)
	SetOS(os string)
	GetComputer() computer.Computer
}

type gamingComputerBuilder struct {
	g computer.Computer
}

// SetProcessor устанавливает процессор
func (b *gamingComputerBuilder) SetProcessor(processor string) {
	b.g.Processor = processor
}

// SetGraphicsCard устанавливает графическую карту
func (b *gamingComputerBuilder) SetGraphicsCard(graphicsCard string) {
	b.g.GraphicsCard = graphicsCard
}

// SetRAM устанавливает оперативную память
func (b *gamingComputerBuilder) SetRAM(ram string) {
	b.g.RAM = ram
}

// SetSSD устанавливает твердотельный накопитель
func (b *gamingComputerBuilder) SetSSD(ssd string) {
	b.g.SSD = ssd
}

// SetOS устанавливает операционную систему
func (b *gamingComputerBuilder) SetOS(os string) {
	b.g.OS = os
}

// GetComputer возвращает экземпляр игрового компьютера
func (b *gamingComputerBuilder) GetComputer() computer.Computer {
	return b.g
}

// NewGamingComputerBuilder возвращает экземпляр gamingComputerBuilder
func NewGamingComputerBuilder() GamingComputerBuilder {
	return &gamingComputerBuilder{}
}
