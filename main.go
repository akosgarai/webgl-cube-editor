package main

import (
	"github.com/hexops/vecty"
)

const (
	PageTitle = "Cube color editor"
)

func main() {
	page := &Page{
		Title:                     PageTitle,
		MeshColor:                 "#ff0000",
		MeshWidth:                 100,
		MeshHeight:                100,
		MeshDepth:                 100,
		BackgroundColor:           "#53c1ff",
		DirectionalLightColor:     "#ffffff",
		DirectionalLightIntensity: 1.0,
		AmbientLightColor:         "#ffffff",
		AmbientLightIntensity:     1.0,
		RotationSpeedY:            200,
		RotationSpeedX:            300,
		SunPosition:               [3]float64{500.0, 256, -256},
	}
	vecty.SetTitle(PageTitle)
	vecty.AddStylesheet("assets/index.css")
	vecty.RenderBody(page)
}
