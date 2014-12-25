package resources

import (
	"github.com/httgo/mock"
	"gopkg.in/nowk/dmx.v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux http.Handler
)

func h(body string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(body))
	})
}

func init() {
	m := dmx.New()
	m.Get("/api/v1/users/current", h(`{"foo":"bar"}`))
	m.Get("/api/v1/summary/daily", h(`{"foo":"bar"}`))
	mux = m.Handler(dmx.NotFound(m))
}

type MyStruct struct {
	Foo string `json:"foo"`
}

func mwaka(t *testing.T) *mock.Mock {
	mo := &mock.Mock{
		Testing: t,
		Ts:      httptest.NewUnstartedServer(mux),
	}
	mo.Start()
	return mo
}
