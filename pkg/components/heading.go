package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// Heading is a general app component.
type Heading struct {
	vecty.Core
	Text string
}

// Render implements vecty.Component for Heading.
func (h *Heading) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Heading1(
			vecty.Text(h.Text),
		),
	)
}
