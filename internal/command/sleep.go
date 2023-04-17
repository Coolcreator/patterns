package command

import (
	"fmt"
	"time"
)

type Sleep struct {
	D time.Duration
}

func (s *Sleep) execute() {
	time.Sleep(s.D)
	fmt.Printf("slept for %.0f seconds\n", s.D.Seconds())
}
