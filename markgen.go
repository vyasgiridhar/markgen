package markgen

const (
	Version = "0.1"
)

func NewMarkgen(port int) *MarkGen {
	return &MarkGen{port, nil, make(chan bool)}
}

type MarkGen struct {
	port       int
	httpServer *MarkdownServer
	stop       chan bool
}
