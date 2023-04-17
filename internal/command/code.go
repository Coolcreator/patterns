package command

import "fmt"

type Code struct {
	Project string
}

func (c *Code) execute() {
	fmt.Printf("coded %s project\n", c.Project)
}
