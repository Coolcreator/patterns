package printer

type printer interface {
	Print()
}

// Print печатает...
func Print(p printer) {
	p.Print()
}
