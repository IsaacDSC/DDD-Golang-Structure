package repositories

import (
	"fmt"
)

var (
	database []User
)

type User struct {
	id         string
	name       string
	currentCNH string
	birthyDay  string
}

type CNH_Repository struct {
	Id               string
	Name             string
	BirthyDay        string
	ValidityCNHTimer string
}

func (i *CNH_Repository) Create(input CNH_Repository) (err error) {
	fmt.Println("Not-Implemented")
	return
}
func (i *CNH_Repository) Update(input CNH_Repository) (err error) {
	fmt.Println("Not-Implemented")
	return
}
func (i *CNH_Repository) Delete(input CNH_Repository) (err error) {
	fmt.Println("Not-Implemented")
	return
}
