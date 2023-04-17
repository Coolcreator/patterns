package cor

import "fmt"

type School struct {
	next Step
}

func (u *School) Execute(p *Person) {
	fmt.Printf("[School] person %s is learning in school\n", p.Name)

	u.next.Execute(p)
}

func (u *School) SetNext(next Step) {
	u.next = next
}
