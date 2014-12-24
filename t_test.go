package resources

import (
	"gopkg.in/nowk/dmx.v2"
	"net/http"
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
	mux = m.Handler(dmx.NotFound(m))
}

type MyStruct struct {
	Foo string `json:"foo"`
}
