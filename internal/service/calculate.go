package service

import (
	"github.com/knudsenTaunus/plan-generator/internal/model"
	"time"
)

type CalculateService struct {
}

func NewCalculate() *CalculateService {
	return &CalculateService{}
}

func (c *CalculateService) CalculatePlan(loanAmount int, nominalRate float32, duration int, startDate time.Time) (*model.Plan, error) {
	return nil, nil
}
