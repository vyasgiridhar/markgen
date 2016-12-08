package markgen

import (
	"fmt"
	"net/http"
	"time"

	ws "github.com/gorilla/websocket"
)

const (
	WriteTimeout = 5 * time.Second
	BufferSize   = 2048
)

var upgrader = ws.Upgrader{
	ReadBufferSize:  BufferSize,
	WriteBufferSize: BufferSize,
}

type Websocket struct {
	watcher *Watcher
}

func NewWebsocket(path string) *Websocket {
	return &Websocket{NewWatch(path)}
}

func (ws *Websocket) Reader(c *ws.Conn, closed chan<- bool) {
	defer c.Close()
	for {
		messageType, _, err := c.NextReader()
		if err != nil || messageType == ws.CloseMessage {
			break
		}
	}
	closed <- true
}

func (ws *Websocket) Writer(c *ws.Conn, closed <-chan bool) {
	ws.watcher.Start()
	defer ws.watcher.Stop()
	defer c.Close()
	for {
		select {
		case data := <-ws.watcher.C.Raw:
			c.SetWriteDeadline(time.Now().Add(WriteTimeout))
			err := c.WriteMessage(ws.TextMessage, MdConverter.Convert(*data))
			if err != nil {
				return
			}
		case <-closed:
			return
		}
	}
}

func (ws *Websocket) Serve(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	sock, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Can't connect to websocket")
		return
	}

	closed := make(chan bool)

	go ws.Reader(sock, closed)
	ws.Writer(sock, closed)
}
