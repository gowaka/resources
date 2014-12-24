package resources

import (
	"github.com/gowaka/api"
)

type CurrentUser struct {
	Client *api.Client
}

func NewCurrentUser(c *api.Client) api.Resource {
	return &CurrentUser{
		Client: c,
	}
}

func (c CurrentUser) Path() string {
	return "/api/v1/users/current"
}

func (c CurrentUser) Get(d interface{}) error {
	return c.Client.Get(&c, d)
}
