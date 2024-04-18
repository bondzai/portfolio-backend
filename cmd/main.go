package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bondzai/test/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://thejb.onrender.com, http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false,
	}))

	handlers.RegisterEndpoints(app)

	app.Listen(":10000")

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	<-gracefulStop

	if err := app.Shutdown(); err != nil {
		fmt.Printf("Error shutting down server %s", err)
	} else {
		fmt.Println("Server gracefully stopped")
	}
}
