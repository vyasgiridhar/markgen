package markgen

import (
	"fmt"

	"github.com/skratchdot/open-golang/open"
)

const (
	Version = "0.1"
)

func NewMarkgen(port int) *MarkGen {
	return &MarkGen{port, nil, make(chan bool)}
}

type MarkGen struct {
	port   int
	Server *MarkdownServer
	stop   chan bool
}

func (m *MarkGen) Run(files ...string) {
	m.Server = NewMarkdownServer(m.port)
	m.Server.Listen()

	for _, file := range files {
		addr := fmt.Sprintf("http://localhost:%d/%s", m.port, file)
		open.Run(addr)
	}

	<-m.stop

}

func (m *MarkGen) Stop() {
	m.Server.Stop()
	m.stop <- true

}
