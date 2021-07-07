package context3

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	data := "Hello World!!"

	t.Run("store returns data", func(t *testing.T) {
		spystore := &SpyStore{response: data, t: t}
		svr := Server(spystore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)

		got := response.Body.String()

		if got != data {
			t.Errorf("want %s, but got %s", data, got)
		}
		//	spystore.assertWasNotCancelled()

	})

	/* 	t.Run("tells store to cancel if request is cancelled", func(t *testing.T) {
	   		spystore := &SpyStore{response: data, t: t}
	   		svr := Server(spystore)

	   		request := httptest.NewRequest(http.MethodGet, "/", nil)

	   		cancellingCtx, cancel := context.WithCancel(request.Context())
	   		request = request.WithContext(cancellingCtx)
	   		response := httptest.NewRecorder()
	   		time.AfterFunc(25*time.Millisecond, cancel)

	   		svr.ServeHTTP(response, request)

	   		spystore.assertWasCancelled()

	   	})
	*/

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(time.Millisecond*5, cancel)
		request = request.WithContext(cancellingCtx)
		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}

	})

}
