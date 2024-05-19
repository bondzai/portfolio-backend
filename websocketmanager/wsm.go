package websocketmanager

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/bondzai/test/interfaces"
	"github.com/gofiber/websocket/v2"
)

type WebSocketManager struct {
	connections []*websocket.Conn
	activeUsers int
	totalUsers  int
	mutex       sync.Mutex
	dbClient    interfaces.MongoDBClientInterface
}

func NewWebSocketManager(dbClient interfaces.MongoDBClientInterface) *WebSocketManager {
	return &WebSocketManager{
		dbClient: dbClient,
	}
}

func (wsm *WebSocketManager) HandleConnection(c *websocket.Conn) {
	wsm.addConnection(c)
	defer wsm.removeConnection(c)

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func (wsm *WebSocketManager) addConnection(c *websocket.Conn) {
	wsm.mutex.Lock()
	defer wsm.mutex.Unlock()

	wsm.connections = append(wsm.connections, c)
	wsm.totalUsers++
	wsm.activeUsers = len(wsm.connections)
	wsm.updateUserCount()
}

func (wsm *WebSocketManager) removeConnection(c *websocket.Conn) {
	wsm.mutex.Lock()
	defer wsm.mutex.Unlock()

	for i, conn := range wsm.connections {
		if conn == c {
			wsm.connections = append(wsm.connections[:i], wsm.connections[i+1:]...)
			break
		}
	}
	wsm.activeUsers = len(wsm.connections)
	wsm.updateUserCount()
}

type UsageCount struct {
	ActiveUsers int `json:"activeUsers"`
	TotalUsers  int `json:"totalUsers"`
}

func (wsm *WebSocketManager) updateUserCount() {
	data := UsageCount{
		ActiveUsers: wsm.activeUsers,
		TotalUsers:  wsm.totalUsers,
	}

	message, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshaling user count:", err)
		return
	}

	for _, conn := range wsm.connections {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Error writing message:", err)
		}
	}
}

func (wsm *WebSocketManager) ResetDailyUserCount() {
	wsm.mutex.Lock()
	defer wsm.mutex.Unlock()

	totalUsers := wsm.totalUsers
	wsm.totalUsers = 0

	wsm.dbClient.SetDataToMongo(&interfaces.User{
		Time:       time.Now(),
		TotalUsers: totalUsers,
	})
}
