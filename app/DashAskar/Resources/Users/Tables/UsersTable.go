package Tables

import "github.com/y-as-7/go-askar/pkg/dash-askar/core"

func UsersTable(table *core.Table) {
	table.Columns = []core.Column{
		{Name: "id", Label: "ID", Sortable: true},
		{Name: "name", Label: "Name", Sortable: true},
		{Name: "email", Label: "Email", Sortable: true},
	}
}
