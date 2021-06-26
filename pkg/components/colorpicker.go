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
}

// Render implements vecty.Component for ColorPicker.
func (c *ColorPicker) Render() vecty.ComponentOrHTML {
	return elem.Input(
		vecty.Markup(
			vecty.Property("id", c.Id),
			prop.Type("color"),
			prop.Value(c.Value),
		),
	)
}
