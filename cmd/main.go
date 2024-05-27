package main

import (
	"log"

	"github.com/bondzai/gogear/toolbox"
	"github.com/bondzai/portfolio-backend/config"
	repository "github.com/bondzai/portfolio-backend/internal/adapters/repository"
	usecases "github.com/bondzai/portfolio-backend/internal/core"
	"github.com/bondzai/portfolio-backend/internal/core/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"

	"github.com/robfig/cron/v3"
)

var cfg = config.LoadConfig()

func main() {
	app := fiber.New()

	toolbox.PPrint(cfg)

	configureCORS(app)

	repo := repository.NewMock()

	mongoClient := initMongoDB()
	userManager := usecases.NewManager(mongoClient)

	setupWebSocketRoutes(app, userManager)
	setupAPIRoutes(app, repo)

	startCronJob(userManager)

	app.Listen(":" + cfg.Port)
}

func initMongoDB() repository.MongoDBClientInterface {
	mongoClient, err := repository.NewMongoDBClient(
		cfg.MongoUrl,
		cfg.MongoDB,
		cfg.MongoCol,
	)

	if err != nil {
		log.Println(err)
	}

	return mongoClient
}

func configureCORS(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowOrigins:     cfg.CorsOrigin,
		AllowHeaders:     cfg.CorsHeader,
		ExposeHeaders:    "Content-Length",
	}))
}

func setupWebSocketRoutes(app *fiber.App, userManager *usecases.Manager) {
	app.Get("/ws", websocket.New(userManager.HandleConnection))
}

func setupAPIRoutes(app *fiber.App, repo *repository.MockRepository) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	app.Get("/certifications", func(c *fiber.Ctx) error {
		certs, _ := services.NewCertService(repo).ReadCerts()
		return c.JSON(certs)
	})

	app.Get("/projects", func(c *fiber.Ctx) error {
		projects, _ := services.NewProjectService(repo).ReadProjects()
		return c.JSON(projects)
	})

	app.Get("/skills", func(c *fiber.Ctx) error {
		skills, _ := services.NewSkillService(repo).ReadSkills()
		return c.JSON(skills)
	})

	app.Get("/wakatime", func(c *fiber.Ctx) error {
		res, _ := services.NewStatService().FetchDataFromAPI()
		return c.JSON(res)
	})
}

func startCronJob(userManager *usecases.Manager) {
	c := cron.New()

	c.AddFunc("59 23 * * *", func() {
		userManager.ResetDailyUserCount()
	})

	c.Start()
	defer c.Stop()

	log.Println("cron started...")
}
