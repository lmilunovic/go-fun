package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := MakeDelayedServer(20 * time.Millisecond)
	fastServer := MakeDelayedServer(0 * time.Millisecond)
	defer slowServer.Close()
	defer fastServer.Close()

	want := fastServer.URL
	got := Racer(slowServer.URL, want)

	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}

}

func MakeDelayedServer(t time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(t)
		w.WriteHeader(http.StatusOK)
	}))
}
