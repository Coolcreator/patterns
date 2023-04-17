package builder

type WorkingComputerBuilder struct {
	Computer
}

func (c *WorkingComputerBuilder) SetProcessor() {
	c.Processor = "Intel Core i7 8th Gen 8665U (1.90GHz)"
}

func (c *WorkingComputerBuilder) SetGraphicsCard() {
	c.GraphicsCard = "Intel UHD Graphics 620"
}

func (c *WorkingComputerBuilder) SetRAM() {
	c.RAM = "8GB DDR4"
}

func (c *WorkingComputerBuilder) SetSSD() {
	c.SSD = "512 GB PCIe SSD"
}

func (c *WorkingComputerBuilder) SetOS() {
	c.OS = "Windows 11 Pro"
}

func (c *WorkingComputerBuilder) GetComputer() Computer {
	return c.Computer
}
