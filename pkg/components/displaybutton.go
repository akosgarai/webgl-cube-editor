package components

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
)

// DispalButton is a general app component.
type DispalButton struct {
	vecty.Core
	Id                 string
	Label              string
	TabulationClass    string
	TargetFormSelector string
	OffIcon            string
	OnIcon             string
}

// Render implements vecty.Component for DispalButton.
func (i *DispalButton) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("row", i.TabulationClass),
		),
		elem.Button(
			vecty.Markup(
				event.Click(func(e *vecty.Event) {
					display := js.Global.Get("document").Call("querySelector", i.TargetFormSelector).Get("style").Get("display").String()
					if display == "none" {
						js.Global.Get("document").Call("querySelector", i.TargetFormSelector).Get("style").Set("display", "block")
						js.Global.Get("document").Call("querySelector", "#"+i.Id).Set("innerText", i.OnIcon)
					} else {
						js.Global.Get("document").Call("querySelector", i.TargetFormSelector).Get("style").Set("display", "none")
						js.Global.Get("document").Call("querySelector", "#"+i.Id).Set("innerText", i.OffIcon)
					}
				}),
			),
			elem.Span(
				vecty.Markup(
					vecty.Class("material-icons"),
					prop.ID(i.Id),
				),
				vecty.Text(i.OffIcon),
			),
			vecty.Text(i.Label),
		),
	)
}
