package main

import (
	"strconv"

	"github.com/akosgarai/webgl-cube-editor/pkg/components"
	"github.com/akosgarai/webgl-cube-editor/pkg/wglrenderer"
	"github.com/divan/three"
	"github.com/gopherjs/gopherjs/js"
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
)

const (
	CubeColorId       = "cube-color"
	CubeWidthId       = "cube-width"
	CubeHeightId      = "cube-height"
	CubeDepthId       = "cube-dept"
	BackgroundColorId = "background-color"
)

// Page is a top-level app component.
type Page struct {
	vecty.Core
	Title           string
	MeshColor       string
	BackgroundColor string
	MeshWidth       int
	MeshHeight      int
	MeshDepth       int
	scene           *three.Scene
	camera          three.PerspectiveCamera
	renderer        *three.WebGLRenderer
	mesh            *three.Mesh
}

// Render implements vecty.Component for Page.
func (p *Page) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(
			event.Change(func(e *vecty.Event) {
				updateGeometry := false
				switch e.Target.Get("id").String() {
				case CubeColorId:
					p.MeshColor = e.Target.Get("value").String()
					// material
					params := three.NewMaterialParameters()
					params.Color = three.NewColor(p.MeshColor)
					p.mesh.Material = three.NewMeshLambertMaterial(params)
					break
				case CubeWidthId:
					p.MeshWidth, _ = strconv.Atoi(e.Target.Get("value").String())
					updateGeometry = true
					break
				case CubeHeightId:
					p.MeshHeight, _ = strconv.Atoi(e.Target.Get("value").String())
					updateGeometry = true
					break
				case CubeDepthId:
					p.MeshDepth, _ = strconv.Atoi(e.Target.Get("value").String())
					updateGeometry = true
					break
				case BackgroundColorId:
					p.BackgroundColor = e.Target.Get("value").String()
					p.scene.Background = three.NewColor(p.BackgroundColor)
					break
				}
				if updateGeometry {
					// size
					p.mesh.Geometry = three.NewBoxGeometry(&three.BoxGeometryParameters{
						Width:  float64(p.MeshWidth),
						Height: float64(p.MeshHeight),
						Depth:  float64(p.MeshDepth),
					})
				}
			}),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("container"),
			),
			elem.Div(
				vecty.Markup(
					prop.ID("form-container"),
					vecty.Class("row"),
				),
				&components.Heading{Text: p.Title},
				&components.ColorPicker{Id: CubeColorId, Value: p.MeshColor, Label: "Cube Color:"},
				&components.ColorPicker{Id: BackgroundColorId, Value: p.BackgroundColor, Label: "Background:"},
				&components.NumericInput{Id: CubeWidthId, Value: p.MeshWidth, Label: "Cube Width:"},
				&components.NumericInput{Id: CubeHeightId, Value: p.MeshHeight, Label: "Cube Height:"},
				&components.NumericInput{Id: CubeDepthId, Value: p.MeshDepth, Label: "Cube Depth:"},
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
		Width:  float64(p.MeshWidth),
		Height: float64(p.MeshHeight),
		Depth:  float64(p.MeshDepth),
	})
	p.mesh = three.NewMesh(geom, mat)
	p.scene.Add(p.mesh)
	p.scene.Background = three.NewColor(p.BackgroundColor)

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
	p.renderer.Render(p.scene, p.camera)
}
