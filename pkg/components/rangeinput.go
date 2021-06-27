package components

import (
	"strconv"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// RangeInput is a general app component.
type RangeInput struct {
	vecty.Core
	Id        string
	Value     int
	MinValue  int
	MaxValue  int
	StepValue int
	Label     string
}

// Render implements vecty.Component for RangeInput.
func (i *RangeInput) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("col-50"),
		),
		&Label{Text: i.Label, For: i.Id},
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
						vecty.Property("id", i.Id),
						vecty.Property("min", i.MinValue),
						vecty.Property("max", i.MaxValue),
						vecty.Property("step", i.StepValue),
						prop.Type("range"),
						prop.Value(strconv.Itoa(i.Value)),
					),
				),
			),
		),
	)
}
