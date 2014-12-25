package resources

import (
	"github.com/gowaka/api"
	"gopkg.in/nowk/assert.v2"
	"testing"
)

func TestCurrentUser(t *testing.T) {
	mo := mwaka(t)
	defer mo.Done()

	waka, err := api.NewClient("...", func(c *api.Client) {
		c.Client = mo
	})
	if err != nil {
		t.Fatal(err)
	}

	var m MyStruct
	err = waka.Get(CurrentUser{}, &m)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "bar", m.Foo)

	reqs := mo.History("GET", "https://wakatime.com/api/v1/users/current")
	assert.Equal(t, 1, len(reqs))
}
