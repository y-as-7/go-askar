package schemas

import "github.com/y-as-7/go-askar/pkg/dash-askar/core"

type Schema struct {
	Components []Component
}

type Component = core.Component
type Form = core.Form
type Table = core.Table
type Infolist = core.Infolist
type Entry = core.Entry
type Column = core.Column

func Make() *Schema {
	return &Schema{}
}

func (s *Schema) Components(components []Component) *Schema {
	s.Components = components
	return s
}

func (s *Schema) GetComponents() []Component {
	return s.Components
}