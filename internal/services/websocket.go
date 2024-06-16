package services

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/bondzai/portfolio-backend/internal/domain"
	"github.com/bondzai/portfolio-backend/internal/repository"
	"github.com/gofiber/websocket/v2"
	"github.com/robfig/cron/v3"
)

type WsService struct {
	connections []*websocket.Conn
	activeUsers int
	totalUsers  int
	mutex       sync.Mutex
	dbClient    repository.MongoDBClientInterface
}

func NewWsService(dbClient repository.MongoDBClientInterface) *WsService {
	return &WsService{
		dbClient: dbClient,
	}
}

func (s *WsService) AddConnection(c *websocket.Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.connections = append(s.connections, c)
	s.totalUsers++
	s.activeUsers = len(s.connections)
	s.updateUserCount()
}

func (s *WsService) RemoveConnection(c *websocket.Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, conn := range s.connections {
		if conn == c {
			s.connections = append(s.connections[:i], s.connections[i+1:]...)
			break
		}
	}
	s.activeUsers = len(s.connections)
	s.updateUserCount()
}

type usageCount struct {
	ActiveUsers int `json:"activeUsers"`
	TotalUsers  int `json:"totalUsers"`
}

func (s *WsService) updateUserCount() {
	data := usageCount{
		ActiveUsers: s.activeUsers,
		TotalUsers:  s.totalUsers,
	}

	message, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshaling user count:", err)
		return
	}

	for _, conn := range s.connections {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Error writing message:", err)
		}
	}
}

func (s *WsService) resetDailyUserCount() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	totalUsers := s.totalUsers
	s.totalUsers = 0

	s.dbClient.InsertOne(
		"usage",
		&domain.TotalUsers{
			Time:       time.Now(),
			TotalUsers: totalUsers,
		},
	)
}

func (s *WsService) StartCronJob() {
	c := cron.New()

	c.AddFunc("59 23 * * *", func() {
		s.resetDailyUserCount()
	})

	c.Start()
	defer c.Stop()

	log.Println("cron started...")
}
