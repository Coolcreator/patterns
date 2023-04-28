package command

// Command представляет собой команду
type Command interface {
	Execute()
	Undo()
}
