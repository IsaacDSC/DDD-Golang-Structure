package shared_app

import (
	"abstract-entities/src/app/usecases/services"
	"abstract-entities/src/infra/repositories"
	"time"

	"github.com/gofrs/uuid"
)

type CNH struct {
	Name      string `json:"name"`
	BirthyDay string `json:"birthyDay"`
}

func GetServiceCNH(body CNH) (output *services.CNH_Service, err error) {
	convertedBirthyDay, err := time.Parse("2006-01-02 15:04:05", body.BirthyDay)
	if err != nil {
		return nil, err
	}
	ID, _ := uuid.NewV4()
	repository := &repositories.CNH_Repository{}
	output = &services.CNH_Service{
		Id:           ID.String(),
		Name:         body.Name,
		BirthyDay:    convertedBirthyDay,
		Repositories: repository,
	}
	return
}
