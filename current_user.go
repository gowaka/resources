package resources

import (
	"github.com/gowaka/api"
)

type CurrentUser struct {
	//
}

func NewCurrentUser() *CurrentUser {
	return &CurrentUser{}
}

func (c CurrentUser) Path() string {
	return "/api/v1/users/current"
}

}
