package controller

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

/*
 * create a room and return code
 */
func CreateRoom(context *gin.Context) {}

/*---------------------------------------------------------------------------*/

/*
 * upgrader to upgrade the GET request to websocket protocol
 * status code 101 switching protocols
 * @todo : add more configuration parameters
 */
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

/*
 * a struct to store websocket connection and a mutex lock to avoid
 * concurrent read and write on the socket
 */
type connection struct {
	userID int64
	conn   *websocket.Conn
	mutex  sync.RWMutex
}

/*
 * map to store connections specific to room (room code -> (user_id -> connection))
 * mutex lock to avoid concurrent read, writes on map
 */
var roomConnections map[string]map[int64]*connection
var roomConnectionsMutex sync.RWMutex

/*
 * join a room with code (create a websocket connection)
 */
func JoinRoom(context *gin.Context) {
	roomCode := context.Param("roomCode")
	c, err := upgrader.Upgrade(context.Writer, context.Request, nil)

	if err != nil {
		context.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error joining chat room",
			"error":   err.Error(),
		})
		return
	}

	// @todo : add user id
	conn := &connection{
		conn: c,
	}

	// update room connections object
	roomConnectionsMutex.Lock()
	if roomConnections[roomCode] == nil {
		roomConnections[roomCode] = make(map[int64]*connection)
	}
	roomConnections[roomCode][conn.userID] = conn
	roomConnectionsMutex.Unlock()

	// handle read message loop
	go handleReadMessages(conn, roomCode)
}

type message struct {
	UserName string `json:"username"`
	Body     string `json:"body"`
}

/*
 * handle messages that are sent on socket
 * @todo : add user details to message object
 */
func handleReadMessages(conn *connection, roomCode string) {
	for {
		var msg struct {
			Body string `json:"body"`
		}
		err := conn.conn.ReadJSON(&msg)

		if err != nil {
			// connections is interrupted/closed
			roomConnectionsMutex.Lock()
			delete(roomConnections[roomCode], conn.userID)
			roomConnectionsMutex.Unlock()
			return
		}

		go broadcastMessage(&message{
			Body: msg.Body,
		}, conn.userID, roomCode)
	}
}

/*
 * broadcast messages to all connections in a chat room
 */
func broadcastMessage(msg *message, userID int64, roomCode string) {
	roomConnectionsMutex.RLock()
	connections := roomConnections[roomCode]
	roomConnectionsMutex.RUnlock()

	for _, conn := range connections {
		if conn.userID != userID {
			conn.mutex.Lock()
			conn.conn.WriteJSON(msg)
			conn.mutex.Unlock()
		}
	}
}
