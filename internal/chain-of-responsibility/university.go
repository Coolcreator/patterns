package cor

import "fmt"

type University struct {
	next Step
}

func (u *University) Execute(p *Person) {
	fmt.Printf("[University] person %s is learning in university\n", p.Name)

	u.next.Execute(p)
}

func (u *University) SetNext(next Step) {
	u.next = next
}
