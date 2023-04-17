package cor

import "fmt"

type Work struct {
	next Step
}

func (u *Work) Execute(p *Person) {
	fmt.Printf("[Work] person %s is working\n", p.Name)

	u.next.Execute(p)
}

func (u *Work) SetNext(next Step) {
	u.next = next
}
