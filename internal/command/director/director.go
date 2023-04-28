package director

// Command представляет собой интерфейс команды
type command interface {
	Execute()
	Undo()
}

// Director исполняет и отменяет команды
type Director interface {
	ExecuteCommand(cmd command)
	UndoLastCommand()
}

type director struct {
	history []command
}

// ExecuteCommand выполняет команду
func (d *director) ExecuteCommand(cmd command) {
	cmd.Execute()
	d.history = append(d.history, cmd)
}

// UndoLastCommand отменяет последнюю команду
func (d *director) UndoLastCommand() {
	if len(d.history) > 0 {
		lastCmd := d.history[len(d.history)-1]
		lastCmd.Undo()
		d.history = d.history[:len(d.history)-1]
	}
}

// NewDirector создает новый экземпляр директора
func NewDirector() Director {
	return &director{}
}
