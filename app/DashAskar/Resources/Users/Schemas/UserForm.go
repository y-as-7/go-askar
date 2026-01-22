package Schemas

import "github.com/y-as-7/go-askar/library/dash-askar/core"

func UserForm(form *core.Form) {
	form.Schema = []core.Component{
		core.TextInput("name", "Name"),
		core.EmailInput("email", "Email"),
	}
}
