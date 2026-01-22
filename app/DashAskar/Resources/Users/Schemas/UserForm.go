package Schemas

import "github.com/y-as-7/go-askar/library/dash-askar/core"

func UserForm(form *core.Form) {
	form.Schema = []core.Component{
		core.TextInput("name", "Name").
			Required().
			MinLength(2).
			MaxLength(100).
			Placeholder("Enter your full name").
			Help("Please enter your first and last name").
			ToComponent(),

		core.TextInput("email", "Email").
			Required().
			Email().
			MaxLength(255).
			Placeholder("user@example.com").
			Help("This will be your login email address").
			ToComponent(),
	}
}
