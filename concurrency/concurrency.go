package concurrency

import (
	"fmt"
	"sync"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)
	for _, url := range urls {
		result[url] = wc(url)
	}

	return result
}

// ConCheckWebsite 并发执行 注意闭包 存在多协程同时写同一个内存地址的问题
func ConCheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	var wg sync.WaitGroup
	result := make(map[string]bool)
	for _, url := range urls {
		url := url
		wg.Add(1)
		go func() {
			result[url] = wc(url)
			fmt.Println(url)
			wg.Done()
		}()
	}
	wg.Wait()

	return result
}

type result struct {
	string
	bool
}

// ConCheckWebsiteWithChannel 用channel 避免多协程同时写同一块内存（数据竞争）
func ConCheckWebsiteWithChannel(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		url := url
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
