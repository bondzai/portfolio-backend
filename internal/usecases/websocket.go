package usecases

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/bondzai/portfolio-backend/config"
	"github.com/bondzai/portfolio-backend/internal/repositories"
	"github.com/gofiber/websocket/v2"
)

type usageCount struct {
	ActiveUsers int `json:"activeUsers"`
	TotalUsers  int `json:"totalUsers"`
}

type (
	WsService interface {
		AddConnection(c *websocket.Conn)
		RemoveConnection(c *websocket.Conn)
	}

	wsService struct {
		connections []*websocket.Conn
		activeUsers int
		totalUsers  int
		mutex       sync.Mutex
		dbClient    repositories.MongoDBClient
		kafkaRepo   repositories.KafkaRepository
	}
)

func NewWsService(dbClient repositories.MongoDBClient, kafkaRepo repositories.KafkaRepository) WsService {
	return &wsService{
		dbClient:  dbClient,
		kafkaRepo: kafkaRepo,
	}
}

func (u *wsService) AddConnection(c *websocket.Conn) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	u.connections = append(u.connections, c)
	u.totalUsers++
	u.activeUsers = len(u.connections)
	u.updateUserCount()

	u.kafkaRepo.Publish(config.AppConfig.KafKaTopic, u.totalUsers)
}

func (u *wsService) RemoveConnection(c *websocket.Conn) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	for i, conn := range u.connections {
		if conn == c {
			u.connections = append(u.connections[:i], u.connections[i+1:]...)
			break
		}
	}
	u.activeUsers = len(u.connections)
	u.updateUserCount()
}

func (u *wsService) updateUserCount() {
	data := usageCount{
		ActiveUsers: u.activeUsers,
		TotalUsers:  u.totalUsers,
	}

	message, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshaling user count:", err)
		return
	}

	for _, conn := range u.connections {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Error writing message:", err)
		}
	}
}
