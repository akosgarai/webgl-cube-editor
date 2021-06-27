package main

import (
	"github.com/hexops/vecty"
)

const (
	PageTitle = "Cube color editor"
)

func main() {
	page := &Page{Title: PageTitle, MeshColor: "#ff0000", MeshWidth: 200, MeshHeight: 200, MeshDepth: 200}
	vecty.SetTitle(PageTitle)
	vecty.AddStylesheet("index.css")
	vecty.RenderBody(page)
}
