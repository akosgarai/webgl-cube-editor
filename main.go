package main

import (
	"github.com/hexops/vecty"
)

const (
	PageTitle = "Cube color editor"
)

func main() {
	page := &Page{
		Title:           PageTitle,
		MeshColor:       "#ff0000",
		MeshWidth:       200,
		MeshHeight:      200,
		MeshDepth:       200,
		BackgroundColor: "#000000",
		LightColor:      "#ffffff",
		RotationSpeedY:  100,
		RotationSpeedX:  0,
	}
	vecty.SetTitle(PageTitle)
	vecty.AddStylesheet("assets/index.css")
	vecty.RenderBody(page)
}
