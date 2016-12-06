package MarkGen

import "time"

const (
	WatcherInterval = 500
	DataChanSize    = 10
)

type DataChannel struct {
	Raw chan *[]byte
	Req chan bool
}

type Watch struct {
	path   string
	ticker *time.Ticker
	stop   chan bool
	C      *DataChannel
}

func newWatch(path string) *Watch {
	dataChan := DataChan{make(chan *[]byte, DataChanSize), make(chan bool)}
	return &Watcher{path, nil, nil, &dataChan}
}
