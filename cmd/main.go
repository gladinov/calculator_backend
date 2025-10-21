package main

import (
	"log"
	calculationservice "main/internal/calculationService"
	"main/internal/db"
	"main/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not init DB:%v", err)
	}
	repo := calculationservice.NewCalculationRepository(db)
	service := calculationservice.NewCalculationService(repo)
	handlers := handlers.NewCalculationHandler(service)
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", handlers.GetCalculations)
	e.POST("/calculations", handlers.PostCalculations)
	e.DELETE("/calculations/:id", handlers.DeleteCalculations)
	e.PATCH("/calculations/:id", handlers.PatchCalculations)

	e.Start("localhost:8080")

}
