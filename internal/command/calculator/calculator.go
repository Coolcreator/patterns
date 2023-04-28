package calculator

// Calculator представляет собой калькулятор
type Calculator interface {
	Add(operand int)
	Subtract(operand int)
	GetValue() int
}

type calculator struct {
	value int
}

// Add складывает значение и операнд
func (c *calculator) Add(operand int) {
	c.value += operand
}

// Subtract вычитает из значения операнд
func (c *calculator) Subtract(operand int) {
	c.value -= operand
}

// GetValue возвращает значение
func (c *calculator) GetValue() int {
	return c.value
}

func NewCalculator() Calculator {
	return &calculator{}
}
