package handlers

import (
	calculationservice "main/internal/calculationService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CalculationHandler interface {
	GetCalculations(c echo.Context) error
	PostCalculations(c echo.Context) error
	DeleteCalculations(c echo.Context) error
	PatchCalculations(c echo.Context) error
}

type CalcHandlers struct {
	service calculationservice.CalculationService
}

func NewCalculationHandler(service calculationservice.CalculationService) CalculationHandler {
	return &CalcHandlers{service: service}
}

func (h *CalcHandlers) GetCalculations(c echo.Context) error {
	calculations, err := h.service.GetAllCalculation()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get calculations"})
	}

	return c.JSON(http.StatusOK, calculations)
}

func (h *CalcHandlers) PostCalculations(c echo.Context) error {
	var req calculationservice.CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	calc, err := h.service.CreateCalculation(req.Expression)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not add calculation"})
	}

	return c.JSON(http.StatusCreated, calc)

}

func (h *CalcHandlers) DeleteCalculations(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeleteCalculation(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete calculation"})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *CalcHandlers) PatchCalculations(c echo.Context) error {
	id := c.Param("id")

	var req calculationservice.CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	calc, err := h.service.UpdateCalculationByID(id, req.Expression)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update calculation"})
	}

	return c.JSON(http.StatusOK, calc)
}
