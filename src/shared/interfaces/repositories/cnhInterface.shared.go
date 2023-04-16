package interfaces

import "abstract-entities/src/infra/repositories"

type CNH_Repository interface {
	Create(input repositories.CNH_Repository) (err error)
	Update(input repositories.CNH_Repository) (err error)
	Delete(input repositories.CNH_Repository) (err error)
}
