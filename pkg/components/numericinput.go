package components

import (
	"strconv"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// NumericInput is a general app component.
type NumericInput struct {
	vecty.Core
	Id    string
	Value int
	Label string
}

// Render implements vecty.Component for Numeric Input.
func (i *NumericInput) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&Label{Text: i.Label, For: i.Id},
		elem.Input(
			vecty.Markup(
				vecty.Property("id", i.Id),
				prop.Type("number"),
				prop.Value(strconv.Itoa(i.Value)),
			),
		),
	)
}
