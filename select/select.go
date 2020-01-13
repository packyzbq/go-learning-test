package _select

import (
	"fmt"
	"net/http"
	"time"
)

// Racer 比较两个url 哪个先返回
func Racer(url1, url2 string) string {

	aDur := measureResponseTime(url1)

	bDur := measureResponseTime(url2)

	if aDur < bDur {
		return url1
	}
	return url2
}

func RacerWithSelect(url1, url2 string) (string, error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(2 * time.Second):
		return "", fmt.Errorf("time out waiting for %s and %s", url1, url2)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	startB := time.Now()
	http.Get(url)
	return time.Since(startB)
}
