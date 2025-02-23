package main

import (
	"log"

	"github.com/bondzai/portfolio-backend/config"
	"github.com/bondzai/portfolio-backend/internal/handlers"
	"github.com/bondzai/portfolio-backend/internal/infrastructures"
	"github.com/bondzai/portfolio-backend/internal/usecases"

	"github.com/gofiber/websocket/v2"
)

func main() {
	mongoClient := infrastructures.NewMongo()

	certService := usecases.NewCertService(mongoClient)
	projectService := usecases.NewProjectService(mongoClient)
	skillService := usecases.NewSkillService(mongoClient)
	wakaService := usecases.NewStatService()
	websocketService := usecases.NewWsService()

	restHandler := handlers.NewHttpHandler(
		certService,
		projectService,
		skillService,
		wakaService,
	)

	websocketHandler := handlers.NewWsHandler(websocketService)

	app := infrastructures.NewFiber()

	app.Get("/", restHandler.HealthCheck)
	app.Get("/certifications", restHandler.GetCerts)
	app.Get("/projects", restHandler.GetProjects)
	app.Get("/skills", restHandler.GetSkills)
	app.Get("/wakatime", restHandler.GetWakaStats)

	app.Get("/ws", websocket.New(websocketHandler.HandleConnection))

	if err := app.Listen(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("Failed to set server %v", err)
	}
}
