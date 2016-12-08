package orange

import (
	"fmt"
	"net/http"
	"time"

	goWs "github.com/gorilla/websocket"
)

const (
	WriteTimeout = 5 * time.Second
	BufferSize   = 2048
)

var upgrader = goWs.Upgrader{
	ReadBufferSize:  BufferSize,
	WriteBufferSize: BufferSize,
}

type Websocket struct {
	watcher *Watcher
}

func NewWebsocket(path string) *Websocket {
	return &Websocket{NewWatcher(path)}
}

func (w *Websocket) Reader(c *goWs.Conn, closed chan<- bool) {
	defer c.Close()
	for {
		messageType, _, err := c.NextReader()
		if err != nil || messageType == goWs.CloseMessage {
			break
		}
	}
	closed <- true
}

func (w *Websocket) Writer(c *goWs.Conn, closed <-chan bool) {
	w.watcher.Start()
	defer w.watcher.Stop()
	defer c.Close()
	for {
		select {
		case data := <-w.watcher.C.Raw:
			c.SetWriteDeadline(time.Now().Add(WriteTimeout))
			err := c.WriteMessage(goWs.TextMessage, MdConverter.Convert(*data))
			if err != nil {
				return
			}
		case <-closed:
			return
		}
	}
}

func (w *Websocket) Serve(w http.ResponseWriter, r *http.Request) {
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

	go w.Reader(sock, closed)
	w.Writer(sock, closed)
}
