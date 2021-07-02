package components

import (
	"fmt"
	"strconv"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// IntRangeInput is a general app component.
type IntRangeInput struct {
	vecty.Core
	Id        string
	Value     int
	MinValue  int
	MaxValue  int
	StepValue int
	Label     string
}

// Render implements vecty.Component for IntRangeInput.
func (i *IntRangeInput) Render() vecty.ComponentOrHTML {
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

// FloatRangeInput is a general app component.
type FloatRangeInput struct {
	vecty.Core
	Id        string
	Value     float64
	MinValue  float64
	MaxValue  float64
	StepValue float64
	Label     string
}

// Render implements vecty.Component for FloatRangeInput.
func (i *FloatRangeInput) Render() vecty.ComponentOrHTML {
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
						prop.Value(fmt.Sprintf("%f", i.Value)),
					),
				),
			),
		),
	)
}
