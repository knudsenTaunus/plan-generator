package model

import "time"

type CalculateRequest struct {
	LoanAmount  int       `json:"loanAmount" validate:"required,gt=0"`
	NominalRate float32   `json:"nominalRate" validate:"required,gt=0"`
	Duration    int       `json:"duration" validate:"required,gt=0"`
	StartDate   time.Time `json:"startDate"`
}

type InputParameters struct {
	LoanAmount  int
	NominalRate float32
	Duration    int
	StartDate   time.Time
}

type Plan struct {
}
