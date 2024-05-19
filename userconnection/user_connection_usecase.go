package userconnection

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/bondzai/test/interfaces"
	"github.com/gofiber/websocket/v2"
)

type UserConnectionManager struct {
	connections []*websocket.Conn
	activeUsers int
	totalUsers  int
	mutex       sync.Mutex
	dbClient    interfaces.MongoDBClientInterface
}

func NewUserConnectionManager(dbClient interfaces.MongoDBClientInterface) *UserConnectionManager {
	return &UserConnectionManager{
		dbClient: dbClient,
	}
}

func (ucm *UserConnectionManager) HandleConnection(c *websocket.Conn) {
	ucm.addConnection(c)
	defer ucm.removeConnection(c)

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func (ucm *UserConnectionManager) addConnection(c *websocket.Conn) {
	ucm.mutex.Lock()
	defer ucm.mutex.Unlock()

	ucm.connections = append(ucm.connections, c)
	ucm.totalUsers++
	ucm.activeUsers = len(ucm.connections)
	ucm.updateUserCount()
}

func (ucm *UserConnectionManager) removeConnection(c *websocket.Conn) {
	ucm.mutex.Lock()
	defer ucm.mutex.Unlock()

	for i, conn := range ucm.connections {
		if conn == c {
			ucm.connections = append(ucm.connections[:i], ucm.connections[i+1:]...)
			break
		}
	}
	ucm.activeUsers = len(ucm.connections)
	ucm.updateUserCount()
}

type UsageCount struct {
	ActiveUsers int `json:"activeUsers"`
	TotalUsers  int `json:"totalUsers"`
}

func (ucm *UserConnectionManager) updateUserCount() {
	data := UsageCount{
		ActiveUsers: ucm.activeUsers,
		TotalUsers:  ucm.totalUsers,
	}

	message, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshaling user count:", err)
		return
	}

	for _, conn := range ucm.connections {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Error writing message:", err)
		}
	}
}

func (ucm *UserConnectionManager) ResetDailyUserCount() {
	ucm.mutex.Lock()
	defer ucm.mutex.Unlock()

	totalUsers := ucm.totalUsers
	ucm.totalUsers = 0

	ucm.dbClient.SetDataToMongo(&interfaces.User{
		Time:       time.Now(),
		TotalUsers: totalUsers,
	})
}
