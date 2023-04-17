package main

import (
	com "patterns/internal/command"
	"time"
)

func main() {
	d := com.Developer{Command: &com.Eat{Food: "pie"}}
	d.Do()
	d.Command = &com.Sleep{D: time.Second}
	d.Do()
	d.Command = &com.Code{Project: "ks8"}
	d.Do()
}
