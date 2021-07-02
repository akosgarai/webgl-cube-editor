package forms

import (
	"github.com/akosgarai/webgl-cube-editor/pkg/components"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// CubeRotation form holds the parameters for
// the rotation of the cube.
type CubeRotation struct {
	vecty.Core
	RotationComponentXId string
	RotationXValue       int
	RotationComponentYId string
	RotationYValue       int
}

// Render implements vecty.Component for CubeRotation.
func (f *CubeRotation) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&components.DisplayButton{
			Id:                 "cube-rotation-lock",
			Label:              "Cube Rotation",
			TabulationClass:    "sub-menu",
			TargetFormSelector: "#cube-rotation-container",
			OffIcon:            "open_in_full",
			OnIcon:             "close_fullscreen",
		},
		elem.Div(
			vecty.Markup(
				vecty.Class("row"),
				prop.ID("cube-rotation-container"),
				vecty.Style("display", "none"),
			),
			&components.IntRangeInput{Id: f.RotationComponentYId, Value: f.RotationYValue, Label: "Y Rotation:", MinValue: -1000, MaxValue: 1000, StepValue: 10},
			&components.IntRangeInput{Id: f.RotationComponentXId, Value: f.RotationXValue, Label: "X Rotation:", MinValue: -1000, MaxValue: 1000, StepValue: 10},
		),
	)
}
