package main

import (
	"github.com/hexops/vecty"
)

const (
	PageTitle = "Cube color editor"
)

func main() {
	page := &Page{Title: PageTitle, MeshColor: "#ff0000"}
	vecty.SetTitle(PageTitle)
	vecty.RenderBody(page)
}
