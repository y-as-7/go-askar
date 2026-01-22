package core

type Resource interface {
	GetModel() string
	GetSlug() string
	GetTitle() string
	GetIcon() string
	
	// Table, Form, and Infolist definitions
	Table(table *Table)
	Form(form *Form)
	Infolist(infolist *Infolist)
}

type Table struct {
	Columns []Column
}

type Column struct {
	Name     string
	Label    string
	Sortable bool
}

type Form struct {
	Schema []Component
}

type Infolist struct {
	Schema []Entry
}

type Entry struct {
	Type  string
	Name  string
	Label string
	Value interface{}
}

func TextEntry(name, label string) Entry {
	return Entry{
		Type:  "text",
		Name:  name,
		Label: label,
	}
}

type Component struct {
	Type        string
	Name        string
	Label       string
	Placeholder string
	Required    bool
	Value       interface{}
}

func TextInput(name, label string) Component {
	return Component{
		Type:  "text",
		Name:  name,
		Label: label,
	}
}

func EmailInput(name, label string) Component {
	return Component{
		Type:  "email",
		Name:  name,
		Label: label,
	}
}
