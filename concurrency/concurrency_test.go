package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWbesiteChecker(url string) bool {
	time.Sleep(time.Millisecond * 20)
	if url == "error" {
		return false
	}
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(mockWbesiteChecker, urls)
	}
}

func BenchmarkConCheckWebsite(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(mockWbesiteChecker, urls)
	}
}

func TestCheckWebsites(t *testing.T) {
	website := []string{
		"http://www.baidu.com",
		"http://www.google.com",
		"error",
	}

	actualResults := ConCheckWebsite(mockWbesiteChecker, website)

	want := len(website)
	got := len(actualResults)

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

	expectedResults := map[string]bool{
		"http://www.baidu.com":  true,
		"http://www.google.com": true,
		"error":                 false,
	}

	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
	}
}

func TestConCheckWebsiteWithChannel(t *testing.T) {
	website := []string{
		"http://www.baidu.com",
		"http://www.google.com",
		"error",
	}

	actualResults := ConCheckWebsiteWithChannel(mockWbesiteChecker, website)

	want := len(website)
	got := len(actualResults)

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

	expectedResults := map[string]bool{
		"http://www.baidu.com":  true,
		"http://www.google.com": true,
		"error":                 false,
	}

	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
	}
}

func BenchmarkConCheckWebsiteWithChannel(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		ConCheckWebsiteWithChannel(mockWbesiteChecker, urls)
	}
}
