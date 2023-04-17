package builder

type GamingComputerBuilder struct {
	Computer
}

func (c *GamingComputerBuilder) SetProcessor() {
	c.Processor = "Intel Core i7 12th Gen 12700F (1.6 GHz)"
}

func (c *GamingComputerBuilder) SetGraphicsCard() {
	c.GraphicsCard = "NVIDIA GeForce RTX 3070 8 GB GDDR6"
}

func (c *GamingComputerBuilder) SetRAM() {
	c.RAM = "16GB DDR4"
}

func (c *GamingComputerBuilder) SetSSD() {
	c.SSD = "512 GB PCIe SSD"
}

func (c *GamingComputerBuilder) SetOS() {
	c.OS = "Windows 11 Home"
}

func (c *GamingComputerBuilder) GetComputer() Computer {
	return c.Computer
}
