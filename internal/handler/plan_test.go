package handler

import (
	"bytes"
	"encoding/json"
	"github.com/knudsenTaunus/plan-generator/internal/handler/mocks"
	"github.com/knudsenTaunus/plan-generator/internal/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlan_GeneratePlan(t *testing.T) {
	requestBody := []byte(`{"loanAmount": "5000","nominalRate": "5.0","duration": 24,"startDate": "2018-01-01T00:00:01Z"}`)
	req, err := http.NewRequest(http.MethodPost, "/rest/api/v1/plan/generate", bytes.NewBuffer(requestBody))
	assert.NoError(t, err)

	var cr model.CalculateRequest
	err = json.Unmarshal(requestBody, &cr)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	mockCalculationService := &mocks.CalculationServiceMock{}
	testPlanHandler := New(mockCalculationService)
	mockInputParameters, err := model.NewInputParametersFromRequest(cr, testPlanHandler.Validator)
	assert.NoError(t, err)

	mockCalculationService.Mock.On("CalculatePlan", mockInputParameters).Return(&Plan{}, nil)
	testPlanHandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

}
