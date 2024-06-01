package services

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/bondzai/portfolio-backend/internal/adapters/repository"
	"github.com/bondzai/portfolio-backend/internal/core/models"
	"github.com/gofiber/websocket/v2"
	"github.com/robfig/cron/v3"
)

type wsService struct {
	connections []*websocket.Conn
	activeUsers int
	totalUsers  int
	mutex       sync.Mutex
	dbClient    repository.MongoDBClientInterface
}

func NewWsService(dbClient repository.MongoDBClientInterface) *wsService {
	return &wsService{
		dbClient: dbClient,
	}
}

func (m *wsService) HandleConnection(c *websocket.Conn) {
	m.addConnection(c)
	defer m.removeConnection(c)

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func (m *wsService) addConnection(c *websocket.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.connections = append(m.connections, c)
	m.totalUsers++
	m.activeUsers = len(m.connections)
	m.updateUserCount()
}

func (m *wsService) removeConnection(c *websocket.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for i, conn := range m.connections {
		if conn == c {
			m.connections = append(m.connections[:i], m.connections[i+1:]...)
			break
		}
	}
	m.activeUsers = len(m.connections)
	m.updateUserCount()
}

type UsageCount struct {
	ActiveUsers int `json:"activeUsers"`
	TotalUsers  int `json:"totalUsers"`
}

func (m *wsService) updateUserCount() {
	data := UsageCount{
		ActiveUsers: m.activeUsers,
		TotalUsers:  m.totalUsers,
	}

	message, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshaling user count:", err)
		return
	}

	for _, conn := range m.connections {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Error writing message:", err)
		}
	}
}

func (m *wsService) ResetDailyUserCount() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	totalUsers := m.totalUsers
	m.totalUsers = 0

	m.dbClient.SetDataToMongo(&models.TotalUsers{
		Time:       time.Now(),
		TotalUsers: totalUsers,
	})
}

func (m *wsService) StartCronJob() {
	c := cron.New()

	c.AddFunc("59 23 * * *", func() {
		m.ResetDailyUserCount()
	})

	c.Start()
	defer c.Stop()

	log.Println("cron started...")
}
