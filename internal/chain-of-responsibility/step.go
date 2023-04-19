package chain_of_responsibility

import "patterns/internal/chain-of-responsibility/api/models"

type Step interface {
	Execute(person *models.Person)
	SetNext(step Step)
}
