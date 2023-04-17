package cor

type Step interface {
	Execute(*Person)
	SetNext(Step)
}
