package resources

import (
	"github.com/gowaka/api"
	"github.com/httgo/mock"
	"gopkg.in/nowk/assert.v2"
	"net/http/httptest"
	"testing"
)

func TestCurrentUser(t *testing.T) {
	mo := &mock.Mock{
		Testing: t,
		Ts:      httptest.NewUnstartedServer(mux),
	}
	mo.Start()
	defer mo.Done()

	c, err := api.NewClient("...", func(c *api.Client) {
		c.Client = mo
	})
	if err != nil {
		t.Fatal(err)
	}

	var m MyStruct
	user := NewCurrentUser(c)
	if err = user.Get(&m); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "bar", m.Foo)

	reqs := mo.History("GET", "https://wakatime.com/api/v1/users/current")
	assert.Equal(t, 1, len(reqs))
}
