package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// Select is a general app component.
type Select struct {
	vecty.Core
	Id string
}

// Select implements vecty.Component for Select.
func (s *Select) Render() vecty.ComponentOrHTML {
	return elem.Select(
		vecty.Markup(
			prop.ID(s.Id),
		),
		elem.Option(
			vecty.Markup(
				prop.Value(""),
			),
			vecty.Text("Choose one"),
		),
	)
}
