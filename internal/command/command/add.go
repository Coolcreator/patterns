package command

type calculator interface {
	Add(operand int)
	Subtract(operand int)
	GetValue() int
}

type addCommand struct {
	calculator calculator
	operand    int
}

// Execute исполняет команду сложения
func (c *addCommand) Execute() {
	c.calculator.Add(c.operand)
}

// Undo отменяет команду сложения
func (c *addCommand) Undo() {
	c.calculator.Subtract(c.operand)
}

// NewAddCommand создает новый экземпляр команды сложения
func NewAddCommand(calculator calculator, operand int) Command {
	return &addCommand{
		calculator: calculator,
		operand:    operand,
	}
}
