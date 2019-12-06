package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

// package main

// import (
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"
// )

// var ErrorOutoftime = errors.New("longer than 10 seconds")

// const timeOut = 1*time.Second

// func TestRacer(t *testing.T) {
// 	t.Run("race test", func(t *testing.T) {
// 		slowServer := makedelayedServer(20 * time.Millisecond)
// 		fastServer := makedelayedServer(0)
// 		defer slowServer.Close()
// 		defer fastServer.Close()
// 		fasturl := fastServer.URL
// 		slowurl := slowServer.URL
// 		want := fasturl
// 		got, err := Racer(slowurl, fasturl)
// 		if err!=nil{
// 			t.Fatal("did not expect an error")
// 		}
// 		if got != want {
// 			t.Errorf("want %q but got %q", want, got)
// 		}
// 	})
// 	t.Run("run out of time", func(t *testing.T) {
// 		serverA := makedelayedServer(11 * time.Second)
// 		serverB := makedelayedServer(12 * time.Second)
// 		defer serverA.Close()
// 		defer serverB.Close()
// 		_, err := Racer(serverA.URL, serverB.URL)
// 		if err != nil {
// 			t.Errorf("Expected an error %v",err)
// 		}
// 	})
// }
// func Racer(fast, slow string) (string, error) {
// 	return Configurableouttime(fast, slow, timeOut)
// }
// func Configurableouttime(fast, slow string, timeout time.Duration) (string, error) {
// 	select {
// 	case <-ping(fast):
// 		return fast, nil
// 	case <-ping(slow):
// 		return slow, nil
// 	case <-time.After(timeout):
// 		return "", ErrorOutoftime
// 	}
// }
// func ping(url string) chan struct{} {
// 	ch := make(chan struct{})
// 	go func() {
// 		http.Get(url)
// 		close(ch)
// 	}()
// 	return ch
// }
// func timerecorder(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }
// func makedelayedServer(t time.Duration) *httptest.Server {
// 	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		time.Sleep(t)
// 		w.WriteHeader(http.StatusOK)
// 	}))
// }
