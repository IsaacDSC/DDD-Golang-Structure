package services

import (
	domain "abstract-entities/src/domain/CNH"
	"abstract-entities/src/infra/repositories"
	shared_domain "abstract-entities/src/shared/domain"
	interfaces "abstract-entities/src/shared/interfaces/repositories"
	"time"
)

type CNH_Service struct {
	Id           string
	Name         string
	BirthyDay    time.Time
	Repositories interfaces.CNH_Repository
}

func (this *CNH_Service) Create() (output shared_domain.BaseEntities[domain.Entity], err []string) {
	output = shared_domain.BaseEntities[domain.Entity]{
		Id: this.Id,
		Props: domain.Entity{
			Name:      this.Name,
			BirthyDay: this.BirthyDay,
		},
	}
	output.Props.ValidatePermissionCNH()
	output.Props.GenerateCurrencyCNH()
	if len(output.Props.Errors) == 0 {
		err := this.Repositories.Create(repositories.CNH_Repository{
			Id:               output.Id,
			Name:             output.Props.Name,
			BirthyDay:        output.Props.BirthyDay.Format(time.RFC3339),
			ValidityCNHTimer: output.Props.CurrentCNH,
		})
		if err != nil {
			output.Props.Errors = append(output.Props.Errors, "GENERIC ERROR: "+err.Error())
		}
	}
	err = output.Props.Errors
	return
}
