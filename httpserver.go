package markgen

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ListeningTestInterval = 500
	MaxListeningTestCount = 10
)

type MarkdownServer struct {
	port     int
	listener net.Listener
}

func NewMarkdownServer(port int) *MarkdownServer {
	return &MarkdownServer{port, nil}
}

func (m *MarkdownServer) Addr() string {
	return ":" + strconv.Itoa(m.port)
}

func (m *MarkdownServer) ListenAndServe() {
	var err error
	server := &http.Server{
		Addr:           m.Addr(),
		Handler:        m,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	m.listener, err = net.Listen("tcp", m.Addr())
	if err != nil {
		panic(err)
	}

	server.Serve(m.listener)
}

func (m *MarkdownServer) Listen() {
	go m.ListenAndServe()

	isListening := make(chan bool)
	go func() {
		result := false
		ticker := time.NewTicker(time.Millisecond * ListeningTestInterval)
		for i := 0; i < MaxListeningTestCount; i++ {
			<-ticker.C
			resp, err := http.Get("http://localhost" + m.Addr() + "/ping")
			if err == nil && resp.StatusCode == 200 {
				result = true
				break
			}
		}
		ticker.Stop()
		isListening <- result
	}()

	if <-isListening {
		fmt.Println("Listening", m.Addr(), "...")
	} else {
		panic("Can't connect to server")
	}
}

func (m *MarkdownServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:] // remove '/'
	if path == "ping" {
		w.Write([]byte("pong"))
	} else if isWebsocketRequest(r) {
		NewWebsocket(path).Serve(w, r)
	} else {
		if strings.HasSuffix(path, ".md") || strings.HasSuffix(path, ".markdown") {
			Template(w, path)
		} else {
			m.ServeStatic(w, path)
		}
	}
}

func (m *MarkdownServer) ServeStatic(w http.ResponseWriter, path string) {
	if stat, err := os.Stat(path); err == nil && stat.Mode().IsRegular() {
		file, _ := os.Open(path)
		io.Copy(w, file)
	}
}

func contains(arr []string, needle string) bool {
	for _, v := range arr {
		if strings.Contains(v, needle) {
			return true
		}
	}
	return false
}

func isWebsocketRequest(r *http.Request) bool {
	upgrade := r.Header["Upgrade"]
	connection := r.Header["Connection"]
	return contains(upgrade, "websocket") && contains(connection, "Upgrade")
}

func (m *MarkdownServer) Stop() {
	m.listener.Close()
}
