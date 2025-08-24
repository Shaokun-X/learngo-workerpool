package worker

import (
	"fmt"
	"time"
)

type DownloadTask struct {
	Name string
	URL  string
}

func (t *DownloadTask) GetName() string {
	return t.Name
}

func (t *DownloadTask) Run() (any, error) {
	fmt.Printf("Downloading %s\n", t.URL)
	// TODO make the http request and save to fs
	time.Sleep(1 * time.Second)
	return true, nil
}
