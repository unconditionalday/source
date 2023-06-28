package checker

import (
	"net/http"
	"time"

	"github.com/SlyMarbo/rss"
)

type RSSChecker struct{}

func NewRSSChecker() RSSChecker {
	return RSSChecker{}
}

func (c RSSChecker) Availability(src string) bool {
	// check rss availability
	if _, err := rss.Fetch(src); err != nil {
		return false
	}

	return true
}

func (c RSSChecker) Latency(src string) (int64, error) {
	n := time.Now()
	req, _ := http.NewRequest("GET", src, nil)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, err
	}

	return time.Since(n).Milliseconds(), nil
}
