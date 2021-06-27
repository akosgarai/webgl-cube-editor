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
	LightColorId      = "light-color"
	RotationSpeedYId  = "rotation-speed-y"
	RotationSpeedXId  = "rotation-speed-x"
)

// Page is a top-level app component.
type Page struct {
	vecty.Core
	Title            string
	MeshColor        string
	BackgroundColor  string
	LightColor       string
	MeshWidth        int
	MeshHeight       int
	MeshDepth        int
	RotationSpeedY   int
	RotationSpeedX   int
	scene            *three.Scene
	camera           three.PerspectiveCamera
	renderer         *three.WebGLRenderer
	mesh             *three.Mesh
	directionalLight *three.DirectionalLight

	canvasWidth  float64
	canvasHeight float64
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
				case LightColorId:
					p.LightColor = e.Target.Get("value").String()
					p.directionalLight.Set("color", three.NewColor(p.LightColor))
					break
				case RotationSpeedYId:
					p.RotationSpeedY, _ = strconv.Atoi(e.Target.Get("value").String())
					break
				case RotationSpeedXId:
					p.RotationSpeedX, _ = strconv.Atoi(e.Target.Get("value").String())
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
				),
				&components.Heading{Text: p.Title},
				elem.Div(
					vecty.Markup(
						vecty.Class("row"),
					),
					elem.Button(
						vecty.Markup(
							event.Click(func(e *vecty.Event) {
								display := js.Global.Get("document").Call("querySelector", "#form-items-container").Get("style").Get("display").String()
								if display == "none" {
									js.Global.Get("document").Call("querySelector", "#form-items-container").Get("style").Set("display", "block")
									js.Global.Get("document").Call("querySelector", "#settings-lock").Set("innerText", "close_fullscreen")
								} else {
									js.Global.Get("document").Call("querySelector", "#form-items-container").Get("style").Set("display", "none")
									js.Global.Get("document").Call("querySelector", "#settings-lock").Set("innerText", "open_in_full")
								}
							}),
						),
						elem.Span(
							vecty.Markup(
								vecty.Class("material-icons"),
								prop.ID("settings-lock"),
							),
							vecty.Text("open_in_full"),
						),
						vecty.Text("Settings"),
					),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("row"),
						prop.ID("form-items-container"),
						vecty.Style("display", "none"),
					),
					&components.ColorPicker{Id: CubeColorId, Value: p.MeshColor, Label: "Cube Color:"},
					&components.ColorPicker{Id: BackgroundColorId, Value: p.BackgroundColor, Label: "Background:"},
					&components.ColorPicker{Id: LightColorId, Value: p.LightColor, Label: "Light:"},
					&components.NumericInput{Id: CubeWidthId, Value: p.MeshWidth, Label: "Cube Width:"},
					&components.NumericInput{Id: CubeHeightId, Value: p.MeshHeight, Label: "Cube Height:"},
					&components.NumericInput{Id: CubeDepthId, Value: p.MeshDepth, Label: "Cube Depth:"},
					&components.RangeInput{Id: RotationSpeedYId, Value: p.RotationSpeedY, Label: "Y Rotation:", MinValue: -1000, MaxValue: 1000, StepValue: 10},
					&components.RangeInput{Id: RotationSpeedXId, Value: p.RotationSpeedX, Label: "X Rotation:", MinValue: -1000, MaxValue: 1000, StepValue: 10},
				),
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

	p.canvasWidth = js.Global.Get("document").Call("querySelector", "#canvas-container").Get("clientWidth").Float()
	p.canvasHeight = js.Global.Get("innerHeight").Float() - js.Global.Get("document").Call("querySelector", "#form-container").Get("clientHeight").Float()*1.1
	devicePixelRatio := js.Global.Get("devicePixelRatio").Float()

	// setup camera and scene
	p.camera = three.NewPerspectiveCamera(70, p.canvasWidth/p.canvasHeight, 1, 1000)
	p.camera.Position.Set(0, 0, 400)
	p.scene = three.NewScene()

	p.renderer.SetPixelRatio(devicePixelRatio)
	p.renderer.SetSize(p.canvasWidth, p.canvasHeight, true)

	// lights
	p.directionalLight = three.NewDirectionalLight(three.NewColor(p.LightColor), 1)
	p.directionalLight.Position.Set(0, 256, 256)
	p.scene.Add(p.directionalLight)

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
	windowWidth := js.Global.Get("document").Call("querySelector", "#canvas-container").Get("clientWidth").Float()
	if windowWidth != p.canvasWidth {
		js.Global.Get("document").Call("querySelector", "#canvas-container canvas").Set("width", windowWidth)
		js.Global.Get("document").Call("querySelector", "#canvas-container canvas").Set("style", "width: 100%")
		p.camera.Aspect = windowWidth / p.canvasHeight
		p.camera.UpdateProjectionMatrix()
		p.renderer.SetSize(windowWidth, p.canvasHeight, false)
		p.canvasWidth = windowWidth
	}
	js.Global.Call("requestAnimationFrame", p.animate)
	p.mesh.Rotation.Set("y", p.mesh.Rotation.Get("y").Float()+0.0001*float64(p.RotationSpeedY))
	p.mesh.Rotation.Set("x", p.mesh.Rotation.Get("x").Float()+0.0001*float64(p.RotationSpeedX))
	p.renderer.Render(p.scene, p.camera)
}
