package main

import (
	"github.com/hexops/vecty"
)

const (
	PageTitle = "Cube color editor"
)

func main() {
	page := &Page{Title: PageTitle, MeshColor: "red"}
	vecty.SetTitle(PageTitle)
	vecty.RenderBody(page)
}
