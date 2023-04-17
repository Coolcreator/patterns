package builder

type Computer struct {
	Processor    string
	GraphicsCard string
	RAM          string
	SSD          string
	OS           string
}

type ComputerBuilder interface {
	SetProcessor()
	SetGraphicsCard()
	SetRAM()
	SetSSD()
	SetOS()
	GetComputer() Computer
}
