package main

import (
	"github.com/hexops/vecty"
)

const (
	PageTitle = "Cube color editor"
)

func main() {
	page := &Page{Title: PageTitle, Message: "This is my message to the world."}
	vecty.SetTitle(PageTitle)
	vecty.RenderBody(page)
}
