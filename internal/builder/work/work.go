package work

import "patterns/internal/builder/api/models"

type workingComputerBuilder struct {
	models.Computer
}

func (c *workingComputerBuilder) SetProcessor() {
	c.Processor = "Intel Core i7 8th Gen 8665U (1.90GHz)"
}

func (c *workingComputerBuilder) SetGraphicsCard() {
	c.GraphicsCard = "Intel UHD Graphics 620"
}

func (c *workingComputerBuilder) SetRAM() {
	c.RAM = "8GB DDR4"
}

func (c *workingComputerBuilder) SetMemory() {
	c.SSD = "512 GB PCIe SSD"
}

func (c *workingComputerBuilder) SetOS() {
	c.OS = "Windows 11 Pro"
}

func (c *workingComputerBuilder) GetComputer() models.Computer {
	return c.Computer
}

func NewWorkingComputerBuilder() *workingComputerBuilder {
	return &workingComputerBuilder{}
}
