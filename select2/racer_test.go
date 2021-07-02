package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("sdfsdf", func(t *testing.T) {
		slowServer := makeDelayedHttpServer(10 * time.Microsecond)
		fastServer := makeDelayedHttpServer(1)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, _ := Racer(slowServer.URL, fastServer.URL)
		if got != want {
			t.Errorf("want %q, but got %q", want, got)
		}
	})

	t.Run("returns error if server does not respond within 10 seconds", func(t *testing.T) {
		serverA := makeDelayedHttpServer(11 * time.Second)
		serverB := makeDelayedHttpServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}

	})

}

func makeDelayedHttpServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}
