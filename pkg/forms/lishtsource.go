package forms

import (
	"github.com/akosgarai/webgl-cube-editor/pkg/components"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// Lightsource form holds the parameters for
// the ambient and directional lightsources.
type Lightsource struct {
	vecty.Core
	AmbientLightColorId         string
	AmbientLightColor           string
	AmbientLightIntensityId     string
	AmbientLightIntensity       float64
	DirectionalLightColorId     string
	DirectionalLightColor       string
	DirectionalLightIntensityId string
	DirectionalLightIntensity   float64
}

// Render implements vecty.Component for Lightsource.
func (f *Lightsource) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&components.DisplayButton{
			Id:                 "lightsources-lock",
			Label:              "Lightsources",
			TabulationClass:    "sub-menu",
			TargetFormSelector: "#lightsources-container",
			OffIcon:            "open_in_full",
			OnIcon:             "close_fullscreen",
		},
		elem.Div(
			vecty.Markup(
				vecty.Class("row"),
				prop.ID("lightsources-container"),
				vecty.Style("display", "none"),
			),
			&components.DisplayButton{
				Id:                 "ambient-lightsources-lock",
				Label:              "Ambient Lightsource",
				TabulationClass:    "sub-menu-2",
				TargetFormSelector: "#ambient-lightsources-container",
				OffIcon:            "open_in_full",
				OnIcon:             "close_fullscreen",
			},
			elem.Div(
				vecty.Markup(
					vecty.Class("row"),
					prop.ID("ambient-lightsources-container"),
					vecty.Style("display", "none"),
				),
				&components.ColorPicker{Id: f.AmbientLightColorId, Value: f.AmbientLightColor, Label: "Light color:"},
				&components.FloatRangeInput{Id: f.AmbientLightIntensityId, Value: f.AmbientLightIntensity, Label: "Intensity:", MinValue: 0, MaxValue: 1, StepValue: 0.01},
			),
			&components.DisplayButton{
				Id:                 "directional-lightsources-lock",
				Label:              "Directional Lightsources",
				TabulationClass:    "sub-menu-2",
				TargetFormSelector: "#directional-lightsources-container",
				OffIcon:            "open_in_full",
				OnIcon:             "close_fullscreen",
			},
			elem.Div(
				vecty.Markup(
					vecty.Class("row"),
					prop.ID("directional-lightsources-container"),
					vecty.Style("display", "none"),
				),
				&components.ColorPicker{Id: f.DirectionalLightColorId, Value: f.DirectionalLightColor, Label: "Light:"},
				&components.FloatRangeInput{Id: f.DirectionalLightIntensityId, Value: f.DirectionalLightIntensity, Label: "Intensity:", MinValue: 0, MaxValue: 1, StepValue: 0.01},
			),
		),
	)
}
