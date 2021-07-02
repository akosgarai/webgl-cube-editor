package forms

import (
	"github.com/akosgarai/webgl-cube-editor/pkg/components"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// Scene form holds the parameters for
// the color of the scene (background color).
type Scene struct {
	vecty.Core
	BackgroundColorId string
	BackgroundColor   string
}

// Render implements vecty.Component for Scene.
func (f *Scene) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&components.DisplayButton{
			Id:                 "scene-lock",
			Label:              "Scene",
			TabulationClass:    "sub-menu",
			TargetFormSelector: "#scene-container",
			OffIcon:            "open_in_full",
			OnIcon:             "close_fullscreen",
		},
		elem.Div(
			vecty.Markup(
				vecty.Class("row"),
				prop.ID("scene-container"),
				vecty.Style("display", "none"),
			),
			&components.ColorPicker{Id: f.BackgroundColorId, Value: f.BackgroundColor, Label: "Background:"},
		),
	)
}
