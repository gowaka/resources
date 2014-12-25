package resources

import (
	"github.com/gowaka/api"
	"net/url"
	"time"
)

type DailySummary struct {
	Start   time.Time
	End     time.Time
	Project string
}

var _ api.Resource = DailySummary{}

func NewDailySummary(opts ...func(*DailySummary)) *DailySummary {
	d := &DailySummary{
		Start: time.Now(),
		End:   time.Now().Add(time.Hour * 24),
	}

	for _, v := range opts {
		v(d)
	}

	return d
}

func (d DailySummary) Path() string {
	return "/api/v1/summary/daily"
}

const (
	datefmt = "01/02/2006"
)

func (d DailySummary) Filter() *url.Values {
	v := &url.Values{
		"start": {d.Start.Format(datefmt)},
		"end":   {d.End.Format(datefmt)},
	}

	if d.Project != "" {
		v.Add("project", d.Project)
	}

	return v
}
