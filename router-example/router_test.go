package router_example

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (StubPlayerStore) GetPlayerScore(name string) int {
	return 0
}

func (StubPlayerStore) RecordWin(name string) {
	return
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []Player

		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", response.Body, err)
		}

		if response.Header().Get("content-type") != "application/json" {
			t.Errorf("response did not have content-type of application/json, got %v", response.Header())
		}

		assertStatus(t, response.Code, http.StatusOK)

		if !reflect.DeepEqual(got, wantedLeague) {
			t.Errorf("got %v want %v", got, wantedLeague)
		}
	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("want %d, but got %d", want, got)
	}
}
