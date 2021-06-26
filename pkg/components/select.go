package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// Select is a general app component.
type Select struct {
	vecty.Core
	Id             string
	Options        map[string]string
	SelectedOption string
}

// Select implements vecty.Component for Select.
func (s *Select) Render() vecty.ComponentOrHTML {
	var options vecty.List
	for k, v := range s.Options {
		selected := false
		if k == s.SelectedOption {
			selected = true
		}
		opt := elem.Option(
			vecty.Markup(
				prop.Value(k),
				vecty.Property("selected", selected),
			),
			vecty.Text(v),
		)
		options = append(options, opt)
	}
	return elem.Select(
		vecty.Markup(
			prop.ID(s.Id),
		),
		options,
	)
}
