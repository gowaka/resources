package resources

import (
	"github.com/gowaka/api"
	"net/url"
)

type CurrentUser struct {
	//
}

var _ api.Resource = CurrentUser{}

func NewCurrentUser() *CurrentUser {
	return &CurrentUser{}
}

func (c CurrentUser) Path() string {
	return "/api/v1/users/current"
}

func (c CurrentUser) Filter() *url.Values {
	return nil
}
