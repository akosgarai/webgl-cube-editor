package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// ColorPicker is a general app component.
type ColorPicker struct {
	vecty.Core
	Id    string
	Value string
	Label string
}

// Render implements vecty.Component for ColorPicker.
func (c *ColorPicker) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("col-50"),
		),
		&Label{Text: c.Label, For: c.Id},
		elem.Div(
			vecty.Markup(
				vecty.Class("col-75"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Style("padding-right", "20px"),
				),
				elem.Input(
					vecty.Markup(
						vecty.Property("id", c.Id),
						prop.Type("color"),
						prop.Value(c.Value),
					),
				),
			),
		),
	)
}
