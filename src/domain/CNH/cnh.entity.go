package domain

import "time"

type Entity struct {
	Name       string
	BirthyDay  time.Time
	CurrentCNH string
	Valid      bool
	Errors     []string
}

func (this *Entity) GenerateCurrencyCNH() {
	if this.Valid != true {
		this.Errors = append(this.Errors, "NOT_PERMITED_YEAR_NOT_VALID")
		return
	}
	currencyDate := time.Now()
	this.CurrentCNH = time.Date(currencyDate.Year()+1, currencyDate.Month(), currencyDate.Day(), 0, 0, 0, 0, time.UTC).Format(time.RFC3339)
	return
}

func (this *Entity) ValidatePermissionCNH() {
	currencyDate := time.Now()
	validYear := (currencyDate.Year() - this.BirthyDay.Year()) >= 18
	validMonth := currencyDate.Month() <= this.BirthyDay.Month()
	if validMonth && validYear {
		this.Valid = true
		return
	}
	this.Valid = false
	return
}
