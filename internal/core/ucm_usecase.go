package usecases

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	repository "github.com/bondzai/portfolio-backend/internal/adapters/repository"
	"github.com/gofiber/websocket/v2"
)

type Manager struct {
	connections []*websocket.Conn
	activeUsers int
	totalUsers  int
	mutex       sync.Mutex
	dbClient    repository.MongoDBClientInterface
}

func NewManager(dbClient repository.MongoDBClientInterface) *Manager {
	return &Manager{
		dbClient: dbClient,
	}
}

func (m *Manager) HandleConnection(c *websocket.Conn) {
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

func (m *Manager) addConnection(c *websocket.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.connections = append(m.connections, c)
	m.totalUsers++
	m.activeUsers = len(m.connections)
	m.updateUserCount()
}

func (m *Manager) removeConnection(c *websocket.Conn) {
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

func (m *Manager) updateUserCount() {
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

func (m *Manager) ResetDailyUserCount() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	totalUsers := m.totalUsers
	m.totalUsers = 0

	m.dbClient.SetDataToMongo(&repository.User{
		Time:       time.Now(),
		TotalUsers: totalUsers,
	})
}
