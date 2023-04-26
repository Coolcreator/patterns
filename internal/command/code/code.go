package code

import "fmt"

// Code представляет собой команду "писать код"
type Code interface {
	Execute()
}

type code struct {
	project string
}

// Execute выполняет команду "писать код"
func (c *code) Execute() {
	fmt.Printf("coded %s project\n", c.project)
}

// NewCode создает новый экземпляр команды "писать код"
func NewCode(project string) Code {
	return &code{project: project}
}
