package calculationservice

import (
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
)

type CalculationService interface {
	CreateCalculation(expression string) (Calculation, error)
	GetAllCalculation() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculationByID(id string, expression string) (Calculation, error)
	DeleteCalculation(id string) error
}

type calcService struct {
	repo CalculationRepository
}

func NewCalculationService(r CalculationRepository) CalculationService {
	return &calcService{repo: r}
}

func (s *calcService) CreateCalculation(expression string) (Calculation, error) {
	result, err := calculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}
	calc := Calculation{
		ID:         uuid.NewString(),
		Expression: expression,
		Result:     result,
	}
	err = s.repo.CreateCalculation(calc)
	if err != nil {
		return Calculation{}, err
	}

	return calc, nil
}

func (s *calcService) DeleteCalculation(id string) error {
	err := s.repo.DeleteCalculation(id)
	return err
}

func (s *calcService) GetAllCalculation() ([]Calculation, error) {
	calculations, err := s.repo.GetAllCalculation()
	if err != nil {
		return nil, err
	}
	return calculations, nil
}

func (s *calcService) GetCalculationByID(id string) (Calculation, error) {
	calc, err := s.repo.GetCalculationByID(id)
	if err != nil {
		return Calculation{}, err
	}
	return calc, nil
}

func (s *calcService) UpdateCalculationByID(id string, expression string) (Calculation, error) {
	// TODO: 2:12:00 https://www.youtube.com/watch?v=gv_EK5kboXc&t=990s
	result, err := calculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}

	calc := Calculation{
		ID:         id,
		Expression: expression,
		Result:     result,
	}

	err = s.repo.UpdateCalculationByID(calc)
	if err != nil {
		return Calculation{}, err
	}

	return calc, nil

}

func calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}
	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result), nil
}
