package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/knudsenTaunus/plan-generator/internal/model"
	"io/ioutil"
	"log"
	"net/http"
)

type CalculationServiceInterface interface {
	CalculatePlan(input *model.InputParameters) (*model.Plan, error)
}

type Plan struct {
	Validator          *validator.Validate
	CalculationService CalculationServiceInterface
}

func (p *Plan) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.GeneratePlan(w, r)
		return
	default:
		http.Error(w, "the requested method is not supported", http.StatusMethodNotAllowed)

	}
}

func New(cs CalculationServiceInterface) *Plan {
	return &Plan{CalculationService: cs, Validator: validator.New()}
}

func (p *Plan) GeneratePlan(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to generate plan", http.StatusInternalServerError)
		return
	}

	var cr model.CalculateRequest
	err = json.Unmarshal(body, &cr)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to generate plan", http.StatusInternalServerError)
	}

	inputParameters, err := model.NewInputParametersFromRequest(cr)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to generate plan", http.StatusInternalServerError)
	}

	err = p.Validator.Struct(inputParameters)
	if err != nil {
		log.Println(err)
		return
	}

	plan, err := p.CalculationService.CalculatePlan(inputParameters)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(plan); err != nil {
		log.Println(err)
		http.Error(w, "handlerErrors parsing results", http.StatusInternalServerError)
		return
	}
}
