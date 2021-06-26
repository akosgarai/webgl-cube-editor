package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// Label is a general app component.
type Label struct {
	vecty.Core
	Text string
	For  string
}

// Render implements vecty.Component for Label.
func (l *Label) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Style("font-family", "monospace"),
			vecty.Class("bold", "label"),
		),
		elem.Label(
			vecty.Markup(
				prop.For(l.For),
			),
			vecty.Text(l.Text),
		),
	)
}
