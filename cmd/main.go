package main

import (
	"flag"
	"log"

	"github.com/bondzai/portfolio-backend/config"
	"github.com/bondzai/portfolio-backend/internal/adapters/handler"
	"github.com/bondzai/portfolio-backend/internal/adapters/repository"
	"github.com/bondzai/portfolio-backend/internal/core/services"
	"github.com/bondzai/portfolio-backend/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

var cfg = config.LoadConfig()

func runSeed() {
	mockRepo := repository.NewMock()
	mongoRepo := initMongoDB()

	certifications, _ := mockRepo.ReadCerts()
	mongoRepo.InsertMany("certifications", utils.ConvertToInterfaceSlice(certifications))
	log.Println("Successfully seeded certifications data to MongoDB")

	projects, _ := mockRepo.ReadProjects()
	mongoRepo.InsertMany("projects", utils.ConvertToInterfaceSlice(projects))
	log.Println("Successfully seeded projects data to MongoDB")

	Skills, _ := mockRepo.ReadSkills()
	mongoRepo.InsertMany("skills", utils.ConvertToInterfaceSlice(Skills))
	log.Println("Successfully seeded Skills data to MongoDB")
}

func runServer() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowOrigins:     cfg.CorsOrigin,
		AllowHeaders:     cfg.CorsHeader,
		ExposeHeaders:    "Content-Length",
	}))

	mongoRepo := initMongoDB()

	certService := services.NewCertService(mongoRepo)
	projectService := services.NewProjectService(mongoRepo)
	skillService := services.NewSkillService(mongoRepo)
	wakaService := services.NewStatService()
	websocketService := services.NewWsService(mongoRepo)

	restHandler := handler.NewHttpHandler(
		certService,
		projectService,
		skillService,
		wakaService,
	)

	websocketHandler := handler.NewWsHandler(websocketService)

	app.Get("/", restHandler.HealthCheck)
	app.Get("/certifications", restHandler.GetCerts)
	app.Get("/projects", restHandler.GetProjects)
	app.Get("/skills", restHandler.GetSkills)
	app.Get("/wakatime", restHandler.GetWakaStats)

	app.Get("/ws", websocket.New(websocketHandler.HandleConnection))

	websocketService.StartCronJob()

	app.Listen(":" + cfg.Port)
}

func main() {
	seedFlag := flag.Bool("seed", false, "Data seeding.")
	flag.Parse()

	if *seedFlag {
		runSeed()
	} else {
		runServer()
	}
}

func initMongoDB() repository.MongoDBClientInterface {
	mongoRepo, err := repository.NewMongoDBClient(
		cfg.MongoUrl,
		cfg.MongoDB,
	)

	if err != nil {
		log.Println(err)
	}

	return mongoRepo
}
