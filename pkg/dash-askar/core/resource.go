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

// FormField interface to support both old and new input types
type FormField interface {
	ToComponent() Component
}

// AddField adds a field to the form (supports both Component and fluent inputs)
func (f *Form) AddField(field interface{}) {
	switch v := field.(type) {
	case Component:
		f.Schema = append(f.Schema, v)
	case *TextInputField:
		f.Schema = append(f.Schema, v.ToComponent())
	case FormField:
		f.Schema = append(f.Schema, v.ToComponent())
	}
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

// Deprecated: Use the fluent TextInput from inputs.go instead
func OldTextInput(name, label string) Component {
	return Component{
		Type:  "text",
		Name:  name,
		Label: label,
	}
}

// Deprecated: Use the fluent EmailInput from inputs.go instead
func OldEmailInput(name, label string) Component {
	return Component{
		Type:  "email",
		Name:  name,
		Label: label,
	}
}
