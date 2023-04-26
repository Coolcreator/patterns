package sleep

import (
	"fmt"
)

// Sleep представляет собой команду "спать"
type Sleep interface {
	Execute()
}

type sleep struct {
	duration int8
}

// Execute выполняет команду "спать"
func (s *sleep) Execute() {
	fmt.Printf("slept for %d minutes\n", s.duration)
}

// NewSleep создает новый экземпляр команды "спать"
func NewSleep(duration int8) Sleep {
	return &sleep{duration: duration}
}
