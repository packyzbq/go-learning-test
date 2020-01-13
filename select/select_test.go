package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := mockServer(20 * time.Millisecond)
	fastServer := mockServer(0)

	defer slowServer.Close()
	defer fastServer.Close()

	want := fastServer.URL
	got := Racer(slowServer.URL, fastServer.URL)

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestRacerWithSelect(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		slowServer := mockServer(20 * time.Millisecond)
		fastServer := mockServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got := Racer(slowServer.URL, fastServer.URL)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := mockServer(3 * time.Second)
		serverB := mockServer(4 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := RacerWithSelect(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func mockServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
