package MarkGen

import (
	"fmt"

	"github.com/skratchdot/open-golang/open"
)

type MarkGen struct {
	port       int
	httpServer *HTTPServer
	stop       chan bool
}

func (*MarkGen) UseBasic() {
	MdConverter.UseBasic()
}

func (m *MarkGen) Run(files ...string) {
	m.httpServer = NewHTTPServer(o.port)
	m.httpServer.Listen()

	for _, file := range files {
		addr := fmt.Sprintf("http://localhost:%d/%s", o.port, file)
		open.Run(addr)
	}

	<-o.stop
}

func (m *MarkGen) Stop() {
	o.httpServer.Stop()
	o.stop <- true
}
