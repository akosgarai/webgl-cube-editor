package forms

import (
	"github.com/akosgarai/webgl-cube-editor/pkg/components"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// CubeDisplay form holds the parameters for
// the display of the cube.
type CubeDisplay struct {
	vecty.Core
	CubeColorId  string
	CubeColor    string
	CubeWidthId  string
	CubeWidth    int
	CubeHeightId string
	CubeHeight   int
	CubeDepthId  string
	CubeDepth    int
}

// Render implements vecty.Component for CubeDisplay.
func (f *CubeDisplay) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&components.DisplayButton{
			Id:                 "cube-display-lock",
			Label:              "Cube Display",
			TabulationClass:    "sub-menu",
			TargetFormSelector: "#cube-display-container",
			OffIcon:            "open_in_full",
			OnIcon:             "close_fullscreen",
		},
		elem.Div(
			vecty.Markup(
				vecty.Class("row"),
				prop.ID("cube-display-container"),
				vecty.Style("display", "none"),
			),
			&components.ColorPicker{Id: f.CubeColorId, Value: f.CubeColor, Label: "Cube Color:"},
			&components.NumericInput{Id: f.CubeWidthId, Value: f.CubeWidth, Label: "Cube Width:"},
			&components.NumericInput{Id: f.CubeHeightId, Value: f.CubeHeight, Label: "Cube Height:"},
			&components.NumericInput{Id: f.CubeDepthId, Value: f.CubeDepth, Label: "Cube Depth:"},
		),
	)
}
