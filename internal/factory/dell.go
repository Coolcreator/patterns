package factory

import "fmt"

type dell struct {
	Name string
}

// DisplayDescription возвращает описание ноутбука
func (d *dell) DisplayDescription() string {
	return fmt.Sprintf("This is a dell laptop named %s.", d.Name)
}
