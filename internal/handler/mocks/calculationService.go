package mocks

import (
	"github.com/knudsenTaunus/plan-generator/internal/model"
	"github.com/stretchr/testify/mock"
)

type CalculationServiceMock struct {
	mock.Mock
}

func (c CalculationServiceMock) CalculatePlan(input *model.InputParameters) (*model.Plan, error) {
	_ = c.Called(input)
	return &model.Plan{}, nil
}
