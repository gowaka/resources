package resources

import (
	"github.com/gowaka/api"
	"gopkg.in/nowk/assert.v2"
	"regexp"
	"testing"
	"time"
)

func TestDailySummary(t *testing.T) {
	mo := mwaka(t)
	defer mo.Done()

	waka, err := api.NewClient("...", func(c *api.Client) {
		c.Client = mo
	})
	if err != nil {
		t.Fatal(err)
	}

	ds := DailySummary{
		Start: time.Date(2014, 12, 14, 0, 0, 0, 0, time.Local),
		End:   time.Date(2014, 12, 25, 0, 0, 0, 0, time.Local),
	}

	var m MyStruct
	err = waka.Get(ds, &m)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "bar", m.Foo)

	reg := regexp.MustCompile(`https://wakatime.com/api/v1/summary/daily`)
	reqs := mo.History("GET", reg)
	assert.Equal(t, 1, len(reqs))

	r := reqs[0]
	v := r.URL.Query()
	assert.Equal(t, "12/14/2014", v.Get("start"))
	assert.Equal(t, "12/25/2014", v.Get("end"))
}

func TestDailySummary_FilterWithProject(t *testing.T) {
	ds := DailySummary{
		Project: "abcdefg",
	}

	v := ds.Filter()
	assert.Equal(t, "abcdefg", v.Get("project"))
}
