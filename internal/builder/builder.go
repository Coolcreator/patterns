package builder

import (
	"patterns/internal/api/computer"
)

type BuildProcess interface {
	SetProcessor() BuildProcess
	SetGraphicsCard() BuildProcess
	SetRAM() BuildProcess
	SetSSD() BuildProcess
	SetOS() BuildProcess
	GetComputer() computer.Computer
}
