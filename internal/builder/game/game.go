package game

import "patterns/internal/builder/api/models"

type gamingComputerBuilder struct {
	models.Computer
}

func (c *gamingComputerBuilder) SetProcessor() {
	c.Processor = "Intel Core i7 12th Gen 12700F (1.6 GHz)"
}

func (c *gamingComputerBuilder) SetGraphicsCard() {
	c.GraphicsCard = "NVIDIA GeForce RTX 3070 8 GB GDDR6"
}

func (c *gamingComputerBuilder) SetRAM() {
	c.RAM = "16GB DDR4"
}

func (c *gamingComputerBuilder) SetMemory() {
	c.SSD = "512 GB PCIe SSD"
}

func (c *gamingComputerBuilder) SetOS() {
	c.OS = "Windows 11 Home"
}

func (c *gamingComputerBuilder) GetComputer() models.Computer {
	return c.Computer
}

func NewGamingComputerBuilder() *gamingComputerBuilder {
	return &gamingComputerBuilder{}
}
