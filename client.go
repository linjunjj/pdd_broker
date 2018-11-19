package main

import (
	"./tool"
	"bytes"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 8192
)

// Client is an middleman between the websocket Client and the hub.ssss
type Client struct {
	hub *Hub
	// The websocket Client.
	ws *websocket.Conn
	// Buffered channel of outbound messages.
	send chan []byte
	// request string
	request map[string]string
	//message
	message []byte
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  8192,
	WriteBufferSize: 8192,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// readPump pumps messages from the websocket Client to the hub.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				tool.Writelog(err.Error())
			}
			break
		}

		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.message = message
		c.hub.broadcast <- c
	}
}

// write writes a message with the given message type and payload.
func (c *Client) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket Client.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.write(websocket.TextMessage, message); err != nil {
				tool.Writelog(err.Error())
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				tool.Writelog(err.Error())
				return
			}
		}
	}
}

func CheckValidate(request map[string]string, addrIp string) bool {
	if request["Time"] == "" {
		return false
	} else {
		timestamp, err := strconv.ParseInt(request["Time"], 10, 64)
		if err != nil {
			tool.Writelog(addrIp + "time convert error")
			return false
		}
		span := timestamp - time.Now().Unix()
		if -300 < span && span < 300 {
			return true
		} else {
			tool.Writelog(addrIp + "token over time 5 minitus")
			return false
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(w http.ResponseWriter, r *http.Request) {
	var pddConn = make(map[string]string)
	for k, v := range r.URL.Query() {
		pddConn[k] = v[0]
	}

	if len(pddConn["access_token"]) <= 0 {
		return
	}

	//list:=strings.Split(pddConn["access_token"],",");
	//if _, ok := clientHub.clients[list[0]]; ok {
	//	clientHub.unregister <- clientHub.clients[pddConn["access_token"]]
	//}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		tool.Writelog(err.Error())
		return
	}
	// login check
	//if CheckValidate(request, ws.RemoteAddr().String()) {

	// init a Client
	c := &Client{hub: clientHub, send: make(chan []byte, 256), ws: conn, request: pddConn}
	c.hub.register <- c

	go c.writePump()
	go c.readPump()
}
