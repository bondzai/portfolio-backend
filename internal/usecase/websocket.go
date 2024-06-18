package usecase

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/bondzai/portfolio-backend/internal/domain"
	"github.com/bondzai/portfolio-backend/internal/repositories"
	"github.com/gofiber/websocket/v2"
	"github.com/robfig/cron/v3"
)

type WsService struct {
	connections []*websocket.Conn
	activeUsers int
	totalUsers  int
	mutex       sync.Mutex
	dbClient    repositories.MongoDBClient
}

func NewWsService(dbClient repositories.MongoDBClient) *WsService {
	return &WsService{
		dbClient: dbClient,
	}
}

func (u *WsService) AddConnection(c *websocket.Conn) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	u.connections = append(u.connections, c)
	u.totalUsers++
	u.activeUsers = len(u.connections)
	u.updateUserCount()
}

func (u *WsService) RemoveConnection(c *websocket.Conn) {
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

type usageCount struct {
	ActiveUsers int `json:"activeUsers"`
	TotalUsers  int `json:"totalUsers"`
}

func (u *WsService) updateUserCount() {
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

func (u *WsService) resetDailyUserCount() {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	totalUsers := u.totalUsers
	u.totalUsers = 0

	u.dbClient.InsertOne(
		"usage",
		&domain.TotalUsers{
			Time:       time.Now(),
			TotalUsers: totalUsers,
		},
	)
}

func (u *WsService) StartCronJob() {
	c := cron.New()

	c.AddFunc("59 23 * * *", func() {
		u.resetDailyUserCount()
	})

	c.Start()
	defer c.Stop()

	log.Println("cron started...")
}
