package markgen

import (
	"fmt"

	"github.com/skratchdot/open-golang/open"
)

const (
	MarkdownChannelSize = 3
	Version             = "0.0.2-dev"
)

type markgen struct {
	port       int
	httpServer *HTTPServer
	stop       chan bool
}

func NewMarkGen(port int) *markgen {
	return &markgen{port, nil, make(chan bool)}
}

func (*markgen) UseBasic() {
	MdConverter.UseBasic()
}

func (m *markgen) Run(files ...string) {
	m.httpServer = NewHTTPServer(m.port)
	m.httpServer.Listen()

	for _, file := range files {
		addr := fmt.Sprintf("http://localhost:%d/%s", m.port, file)
		open.Run(addr)
	}

	<-m.stop
}

func (m *markgen) Stop() {
	m.httpServer.Stop()
	m.stop <- true
}
