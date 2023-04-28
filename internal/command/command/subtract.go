package command

type subtractCommand struct {
	calculator calculator
	operand    int
}

// Execute исполняет команду вычитания
func (c *subtractCommand) Execute() {
	c.calculator.Subtract(c.operand)
}

// Undo отменяет команду вычитания
func (c *subtractCommand) Undo() {
	c.calculator.Add(c.operand)
}

// NewSubtractCommand создает новый экземпляр команды вычитания
func NewSubtractCommand(calculator calculator, operand int) Command {
	return &subtractCommand{
		calculator: calculator,
		operand:    operand,
	}
}
