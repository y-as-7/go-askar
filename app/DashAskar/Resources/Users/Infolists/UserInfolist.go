package Infolists

import "github.com/y-as-7/go-askar/library/dash-askar/core"

func UserInfolist(infolist *core.Infolist) {
	infolist.Schema = []core.Entry{
		core.TextEntry("id", "ID"),
		core.TextEntry("name", "Name"),
		core.TextEntry("email", "Email"),
		core.TextEntry("role", "Role"),
	}
}
