package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
)

// TextArea is a general app component.
type TextArea struct {
	vecty.Core
	Message string
	Id      string
}

// Render implements vecty.Component for TextArea.
func (t *TextArea) Render() vecty.ComponentOrHTML {
	return elem.TextArea(
		vecty.Markup(
			vecty.Style("font-family", "monospace"),
			vecty.Property("rows", 14),
			vecty.Property("cols", 70),
			vecty.Property("id", t.Id),

			// When input is typed into the textarea, update the local
			// component state and rerender.
			event.Input(func(e *vecty.Event) {
				t.Message = e.Target.Get("value").String()
				vecty.Rerender(t)
			}),
		),
		vecty.Text(t.Message),
	)
}
