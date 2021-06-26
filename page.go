package main

import (
	"github.com/akosgarai/webgl-cube-editor/pkg/components"
	"github.com/akosgarai/webgl-cube-editor/pkg/wglrenderer"
	"github.com/divan/three"
	"github.com/gopherjs/gopherjs/js"
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
)

// Page is a top-level app component.
type Page struct {
	vecty.Core
	Title     string
	MeshColor string
	scene     *three.Scene
	camera    three.PerspectiveCamera
	renderer  *three.WebGLRenderer
	mesh      *three.Mesh
}

// Render implements vecty.Component for Page.
func (p *Page) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(
			event.Change(func(e *vecty.Event) {
				p.MeshColor = e.Target.Get("value").String()
			}),
		),
		elem.Div(
			vecty.Markup(
				prop.ID("form-container"),
			),
			&components.Heading{Text: p.Title},
			&components.Label{Text: "Color:", For: "cube-color"},
			&components.ColorPicker{Id: "cube-color", Value: p.MeshColor},
		),
		elem.Div(
			vecty.Markup(
				prop.ID("canvas-container"),
				vecty.Style("width", "90%"),
				vecty.Style("margin-left", "auto"),
				vecty.Style("margin-right", "auto"),
			),
			wglrenderer.WebGLRenderer(wglrenderer.WebGLOptions{
				Init:     p.init,
				Shutdown: p.shutdown,
			}),
		),
	)
}
func (p *Page) init(renderer *three.WebGLRenderer) {
	p.renderer = renderer

	windowWidth := js.Global.Get("document").Call("querySelector", "#canvas-container").Get("clientWidth").Float()
	windowHeight := js.Global.Get("innerHeight").Float() - js.Global.Get("document").Call("querySelector", "#form-container").Get("clientHeight").Float()*2
	devicePixelRatio := js.Global.Get("devicePixelRatio").Float()

	// setup camera and scene
	p.camera = three.NewPerspectiveCamera(70, windowWidth/windowHeight, 1, 1000)
	p.camera.Position.Set(0, 0, 400)
	p.scene = three.NewScene()

	p.renderer.SetPixelRatio(devicePixelRatio)
	p.renderer.SetSize(windowWidth, windowHeight, true)

	// lights
	light := three.NewDirectionalLight(three.NewColor("white"), 1)
	light.Position.Set(0, 256, 256)
	p.scene.Add(light)

	// material
	params := three.NewMaterialParameters()
	params.Color = three.NewColor(p.MeshColor)
	mat := three.NewMeshLambertMaterial(params)

	// cube object
	geom := three.NewBoxGeometry(&three.BoxGeometryParameters{
		Width:  200,
		Height: 200,
		Depth:  200,
	})
	p.mesh = three.NewMesh(geom, mat)
	p.scene.Add(p.mesh)

	// start animation
	p.animate()
}
func (p *Page) shutdown(renderer *three.WebGLRenderer) {
	// After shutdown, we shouldn't use any of these anymore.
	p.scene = nil
	p.camera = three.PerspectiveCamera{}
	p.renderer = nil
	p.mesh = nil
}
func (p *Page) animate() {
	if p.renderer == nil {
		// We shutdown, stop animation.
		return
	}
	js.Global.Call("requestAnimationFrame", p.animate)
	currentRotation := p.mesh.Rotation.Get("y").Float()
	p.mesh.Rotation.Set("y", currentRotation+0.01)
	// material
	params := three.NewMaterialParameters()
	params.Color = three.NewColor(p.MeshColor)
	p.mesh.Material = three.NewMeshLambertMaterial(params)
	p.renderer.Render(p.scene, p.camera)
}
