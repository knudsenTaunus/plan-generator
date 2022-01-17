package model

import (
	"strconv"
	"time"
)

type CalculateRequest struct {
	LoanAmount  string    `json:"loanAmount"`
	NominalRate string    `json:"nominalRate"`
	Duration    int       `json:"duration"`
	StartDate   time.Time `json:"startDate"`
}

type InputParameters struct {
	LoanAmount  float64   `validate:"required,gt=0"`
	NominalRate float64   `validate:"required,gt=0"`
	Duration    int       `validate:"required,gt=0"`
	StartDate   time.Time `validate:"required"`
}

func NewInputParametersFromRequest(cr CalculateRequest) (*InputParameters, error) {
	loanAmount, err := strconv.ParseFloat(cr.LoanAmount, 32)
	if err != nil {
		return nil, err
	}
	nominalRate, err := strconv.ParseFloat(cr.NominalRate, 32)
	if err != nil {
		return nil, err
	}

	return &InputParameters{
		LoanAmount:  loanAmount,
		NominalRate: nominalRate / 100,
		Duration:    cr.Duration,
		StartDate:   cr.StartDate,
	}, nil

}

type BorrowerPayment struct {
	Date                          time.Time `json:"date"`
	BorrowerPaymentAmount         string    `json:"borrowerPaymentAmount"`
	Principal                     string    `json:"principal"`
	Interest                      string    `json:"interest"`
	InitialOutstandingPrincipal   string    `json:"initialOutstandingPrincipal"`
	RemainingOutstandingPrincipal string    `json:"remainingOutstandingPrincipal"`
}

type Plan struct {
	BorrowerPayments []BorrowerPayment `json:"borrowerPayments"`
}
