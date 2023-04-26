package developer

type command interface {
	Execute()
}

// Developer представляет собой разработчика
type Developer interface {
	Do()
	SetCommand(c command) Developer
}

type developer struct {
	command command
}

// Do выполняет команду
func (d *developer) Do() {
	d.command.Execute()
}

// SetCommand устанавливает команду
func (d *developer) SetCommand(command command) Developer {
	d.command = command
	return d
}

// NewDeveloper создает новый экземпляр разработчика
func NewDeveloper() Developer {
	return &developer{}
}
