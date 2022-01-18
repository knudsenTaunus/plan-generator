package service

import (
	"fmt"
	"github.com/knudsenTaunus/plan-generator/internal/model"
	"math"
	"time"
)

const (
	MONTH = 30
	YEAR  = 360
)

type CalculateService struct {
}

func NewCalculate() *CalculateService {
	return &CalculateService{}
}

func (c *CalculateService) CalculatePlan(input *model.InputParameters) (*model.Plan, error) {
	plan := &model.Plan{BorrowerPayments: []model.BorrowerPayment{}}

	remainingOutstandingPrincipal := input.LoanAmount
	date := input.StartDate
	duration := input.Duration

	for remainingOutstandingPrincipal > 0.9 && duration != 0 {
		bpa := c.CalculateAnnuity(input.NominalRate, remainingOutstandingPrincipal, duration)
		interest := c.CalculateInterest(input.NominalRate, remainingOutstandingPrincipal)
		principal := c.CalculatePrincipal(bpa, interest)
		newRemainingOutstandingPrincipal := remainingOutstandingPrincipal - principal

		if remainingOutstandingPrincipal <= bpa {
			newRemainingOutstandingPrincipal = 0
			bpa = newRemainingOutstandingPrincipal
		}

		pr := model.BorrowerPayment{
			Date:                          date,
			BorrowerPaymentAmount:         fmt.Sprintf("%.2f", bpa),
			Principal:                     fmt.Sprintf("%.2f", principal),
			Interest:                      fmt.Sprintf("%.2f", interest),
			InitialOutstandingPrincipal:   fmt.Sprintf("%.2f", remainingOutstandingPrincipal),
			RemainingOutstandingPrincipal: fmt.Sprintf("%.2f", newRemainingOutstandingPrincipal),
		}

		duration--
		date = calculateNextDate(date)
		//date = time.Date(date.Year(), date.Month()+1, getLastDay(date.Month()+1, date.Year()), 0, 0, 0, 0, time.UTC)
		remainingOutstandingPrincipal = newRemainingOutstandingPrincipal
		plan.BorrowerPayments = append(plan.BorrowerPayments, pr)
	}

	return plan, nil
}

func calculateNextDate(date time.Time) time.Time {
	if isLastDay(date.Day(), date.Month(), date.Year()) {
		return time.Date(date.Year(), date.Month()+1, getLastDay(date.Month()+1, date.Year()), 0, 0, 0, 0, time.UTC)
	}
	return date.AddDate(0, 1, 0)

}

func isLastDay(day int, month time.Month, year int) bool {
	if day == 31 {
		return true
	}

	switch month {
	case time.April, time.June, time.September, time.November:
		return day == 30
	case time.February:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) { // leap year
			return day == 29
		}
		return day == 28
	default:
		return day == 30
	}
}

func getLastDay(month time.Month, year int) int {
	switch month {
	case time.February:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) { // leap year
			return 29
		}
		return 28
	default:
		return 30
	}
}

func EndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

func (c *CalculateService) CalculateInterest(rate, outstandingInitialPrincipal float64) float64 {
	return math.Floor(rate*MONTH*outstandingInitialPrincipal) / YEAR * 100 / 100
	//return (rate * MONTH * outstandingInitialPrincipal) / YEAR
}

func (c *CalculateService) CalculatePrincipal(annuity, interest float64) float64 {
	return math.Floor((annuity-interest)*100) / 100
	//return annuity - interest
}

func (c *CalculateService) CalculateBorrowPaymentAmount(principal float64, interest float64) float64 {
	return math.Floor((principal+interest)*100) / 100
	//return principal + interest
}

func (c *CalculateService) CalculateAnnuity(rate, outstandingInitialPrincipal float64, duration int) float64 {

	monthlyRate := ((rate / 12) * 100) / 100

	result := (outstandingInitialPrincipal * monthlyRate) / (1 - math.Pow(1+monthlyRate, float64(-duration)))
	return math.Ceil(result*100) / 100
	//return result
}
