package Users

import (
	"github.com/y-as-7/go-askar/app/DashAskar/Resources/Users/Infolists"
	"github.com/y-as-7/go-askar/app/DashAskar/Resources/Users/Schemas"
	"github.com/y-as-7/go-askar/app/DashAskar/Resources/Users/Tables"
	"github.com/y-as-7/go-askar/library/dash-askar/core"
)

type UserResource struct{}

func (u *UserResource) GetModel() string { return "User" }
func (u *UserResource) GetSlug() string  { return "users" }
func (u *UserResource) GetTitle() string { return "Users" }
func (u *UserResource) GetIcon() string  { return "👥" }

func (u *UserResource) Table(table *core.Table) {
	Tables.UsersTable(table)
}

func (u *UserResource) Form(form *core.Form) {
	Schemas.UserForm(form)
}

func (u *UserResource) Infolist(infolist *core.Infolist) {
	Infolists.UserInfolist(infolist)
}
