package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"github.com/bondzai/test/handlers"
	"github.com/bondzai/test/interfaces"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"github.com/robfig/cron/v3"
)

type WebSocketHandler interface {
	HandleConnection(*websocket.Conn)
}

type UserCountUpdater interface {
	UpdateUserCount()
}

type WebSocketManager struct {
	Connections []*websocket.Conn
	ActiveUsers int
	TotalUsers  int
	mutex       sync.Mutex
	DBClient    interfaces.MongoDBClientInterface
}

func (wsm *WebSocketManager) HandleConnection(c *websocket.Conn) {
	wsm.mutex.Lock()
	wsm.Connections = append(wsm.Connections, c)
	wsm.TotalUsers++
	wsm.ActiveUsers = len(wsm.Connections)
	wsm.mutex.Unlock()
	wsm.UpdateUserCount()

	defer func() {
		wsm.mutex.Lock()
		for i, conn := range wsm.Connections {
			if conn == c {
				wsm.Connections = append(wsm.Connections[:i], wsm.Connections[i+1:]...)
				break
			}
		}
		wsm.ActiveUsers = len(wsm.Connections)
		wsm.mutex.Unlock()
		wsm.UpdateUserCount()
	}()

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			break
		}
		log.Println(err)
	}
}

type UsageCount struct {
	ActiveUsers int `json:"activeUsers"`
	TotalUsers  int `json:"totalUsers"`
}

func (wsm *WebSocketManager) UpdateUserCount() {
	wsm.mutex.Lock()
	defer wsm.mutex.Unlock()

	data := UsageCount{
		ActiveUsers: wsm.ActiveUsers,
		TotalUsers:  wsm.TotalUsers,
	}

	message, err := json.Marshal(data)
	if err != nil {
		return
	}

	for _, conn := range wsm.Connections {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			return
		}
	}
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("GO_CORS_ORIGINS"),
		AllowHeaders:     os.Getenv("GO_CORS_HEADERS"),
		ExposeHeaders:    "Content-Length",
		AllowCredentials: false,
	}))

	mongoClient, err := interfaces.NewMongoDBClient(os.Getenv("GO_MONGODB_URL"), "portfolio", "usage")
	if err != nil {
		log.Print(err)
	}

	wsm := &WebSocketManager{
		DBClient: mongoClient,
	}
	app.Get("/ws", websocket.New(wsm.HandleConnection))

	handlers.RegisterEndpoints(app)

	c := cron.New()
	c.AddFunc("59 23 * * *", func() {
		wsm.mutex.Lock()
		defer wsm.mutex.Unlock()

		totalUsers := wsm.TotalUsers
		wsm.TotalUsers = 0

		wsm.DBClient.SetDataToMongo(&interfaces.User{
			Time:       time.Now(),
			TotalUsers: totalUsers,
		})
	})
	c.Start()
	defer c.Stop()

	log.Println("cron start...")

	app.Listen(":" + os.Getenv("GO_PORT"))
}
