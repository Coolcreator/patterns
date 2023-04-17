package cor

import "fmt"

type Asylum struct {
	next Step
}

func (u *Asylum) Execute(p *Person) {
	fmt.Printf("[Asylum] person %s is in mental institution\n", p.Name)
}

func (u *Asylum) SetNext(next Step) {
	u.next = next
}
