package command

type Developer struct {
	Command Command
}

func (d *Developer) Do() {
	d.Command.execute()
}
