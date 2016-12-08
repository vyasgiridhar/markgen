package markgen

import (
	"io/ioutil"
	"os"
	"time"
)

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

func (w *Watch) Start() {
	go func() {
		w.ticker = time.NewTicker(time.Millisecond * WatcherInterval)
		defer w.ticker.Stop()
		w.stop = make(chan bool)
		var currentTimestamp int64
		for {
			select {
			case <-w.stop:
				return
			case <-w.ticker.C:
				reload := false
				select {
				case <-w.C.Req:
					reload = true
				default:
				}

				info, err := os.Stat(w.path)
				if err != nil {
					continue
				}

				timestamp := info.ModTime().Unix()
				if currentTimestamp < timestamp || reload {
					currentTimestamp = timestamp

					raw, err := ioutil.ReadFile(w.path)
					if err != nil {
						continue
					}

					w.C.Raw <- &raw
				}
			}
		}
	}()

}

func (w *Watch) Stop() {
	w.stop <- true

}
