package models

import (
	"regexp"

	"github.com/ln0rd/api-golang-persons/errs"
)

var (
	nameRegex                   = regexp.MustCompile(`^[a-z A-Z\s]{2,50}$`)
	historyRegex                = regexp.MustCompile(`^[\w\s.,!?'-]{10,500}$`)
	Personalities []Personality = []Personality{}
)

type Personality struct {
	ID      int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name    string `json:"name" gorm:"column:name"`
	History string `json:"history" gorm:"column:history"`
}

func (Personality) TableName() string {
	return "personalities"
}

func (p *Personality) Validate() error {
	if !nameRegex.MatchString(p.Name) {
		return errs.NewValidateRequestError("invalid name format: must be between 2-50 characters and contain only letters and spaces")
	}

	if !historyRegex.MatchString(p.History) {
		return errs.NewValidateRequestError("invalid history format: must be between 10-500 characters and contain only letters, numbers, spaces and basic punctuation")
	}

	return nil
}
