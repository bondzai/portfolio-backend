package main

import (
	"flag"
	"log"

	"github.com/bondzai/portfolio-backend/config"
	repository "github.com/bondzai/portfolio-backend/internal/adapters/repository"
	usecases "github.com/bondzai/portfolio-backend/internal/core"
	"github.com/bondzai/portfolio-backend/internal/core/models"
	"github.com/bondzai/portfolio-backend/internal/core/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"

	"github.com/robfig/cron/v3"
)

var conf = config.GetConfig()

func main() {
	startCronFlag := flag.Bool("cron", false, "Start Cronjob flag")
	flag.Parse()

	app := fiber.New()

	configureCORS(app)

	repo := repository.NewMock()

	mongoClient := initMongoDB()
	userManager := usecases.NewManager(mongoClient)

	setupWebSocketRoutes(app, userManager)
	setupAPIRoutes(app, repo)

	if *startCronFlag {
		startCronJob(userManager)
	}

	app.Listen(":" + conf.Port)
}

func initMongoDB() repository.MongoDBClientInterface {
	mongoClient, err := repository.NewMongoDBClient(
		conf.MongoUrl,
		conf.MongoDB,
		conf.MongoCol,
	)

	if err != nil {
		log.Println(err)
	}

	return mongoClient
}

func configureCORS(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     conf.CorsOrigin,
		AllowHeaders:     conf.CorsHeader,
		ExposeHeaders:    "Content-Length",
		AllowCredentials: false,
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
		return c.JSON(models.Wakatime)
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
