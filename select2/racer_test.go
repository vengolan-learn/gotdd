package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares speeds of serves, returning the url of the fastest response", func(t *testing.T) {
		slowServer := makeDelayedHttpServer(10 * time.Millisecond)
		fastServer := makeDelayedHttpServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, err := Racer(slowServer.URL, fastServer.URL)

		if err != nil {
			t.Fatalf("did not expect an error, but got one %v", err)
		}
		if got != want {
			t.Errorf("want %q, but got %q", want, got)
		}
	})

	t.Run("returns error if server does not respond within 10 seconds", func(t *testing.T) {
		server := makeDelayedHttpServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

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
