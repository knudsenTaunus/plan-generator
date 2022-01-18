package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/knudsenTaunus/plan-generator/internal/model"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestCalculateService_CalculatePlanTrickyDate(t *testing.T) {

	calculateRequest := model.CalculateRequest{
		LoanAmount:  "6000",
		NominalRate: "7",
		Duration:    24,
		StartDate:   time.Date(2022, time.January, 30, 0, 0, 0, 0, time.UTC),
	}

	inputParameters, err := model.NewInputParametersFromRequest(calculateRequest, validator.New())
	assert.NoError(t, err)

	calculateService := NewCalculate()

	plan, err := calculateService.CalculatePlan(inputParameters)
	assert.NoError(t, err)
	log.Println(plan.BorrowerPayments[0].Interest)
	assert.Equal(t, "268.64", plan.BorrowerPayments[0].BorrowerPaymentAmount)
	assert.Equal(t, 28, plan.BorrowerPayments[1].Date.Day())

	assert.Equal(t, calculateRequest.Duration, len(plan.BorrowerPayments))
}

func TestCalculateService_CalculatePlanNormalDate(t *testing.T) {

	calculateRequest := model.CalculateRequest{
		LoanAmount:  "6000",
		NominalRate: "7",
		Duration:    24,
		StartDate:   time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
	}

	inputParameters, err := model.NewInputParametersFromRequest(calculateRequest, validator.New())
	assert.NoError(t, err)

	calculateService := NewCalculate()

	plan, err := calculateService.CalculatePlan(inputParameters)
	assert.NoError(t, err)
	log.Println(plan.BorrowerPayments[0].Interest)
	assert.Equal(t, "268.64", plan.BorrowerPayments[0].BorrowerPaymentAmount)
	assert.Equal(t, 15, plan.BorrowerPayments[1].Date.Day())

	assert.Equal(t, calculateRequest.Duration, len(plan.BorrowerPayments))
}

func TestCalculateService_CalculatePlanNegativeRate(t *testing.T) {

	calculateRequest := model.CalculateRequest{
		LoanAmount:  "6000",
		NominalRate: "-7",
		Duration:    24,
		StartDate:   time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
	}

	inputParameters, err := model.NewInputParametersFromRequest(calculateRequest, validator.New())
	assert.Error(t, err)
	assert.Nil(t, inputParameters)
}
