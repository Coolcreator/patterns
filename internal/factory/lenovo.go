package factory

import "fmt"

type lenovo struct {
	Name string
}

// DisplayDescription возвращает описание ноутбука
func (l *lenovo) DisplayDescription() string {
	return fmt.Sprintf("This is a lenovo laptop named %s.", l.Name)
}
