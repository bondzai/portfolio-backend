package main

import (
	"os"
	"strconv"
	"sync"

	"github.com/bondzai/test/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

var (
	activeUsers int
	mutex       sync.Mutex
	connections []*websocket.Conn
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("GO_CORS_ORIGINS"),
		AllowHeaders:     os.Getenv("GO_CORS_HEADERS"),
		AllowCredentials: false,
	}))

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		mutex.Lock()
		connections = append(connections, c)
		activeUsers = len(connections)
		mutex.Unlock()

		sendActiveUserCount()

		defer func() {
			mutex.Lock()
			for i, conn := range connections {
				if conn == c {
					connections = append(connections[:i], connections[i+1:]...)
					break
				}
			}
			activeUsers = len(connections)
			mutex.Unlock()
			sendActiveUserCount()
		}()

		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				break
			}
			//handle received messages here
		}
	}))

	handlers.RegisterEndpoints(app)

	app.Listen(":" + os.Getenv("GO_PORT"))
}

func sendActiveUserCount() {
	mutex.Lock()
	defer mutex.Unlock()
	for _, conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(activeUsers))); err != nil {
			return
		}
	}
}
