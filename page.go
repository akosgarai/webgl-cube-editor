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
	CubeColorId             = "cube-color"
	CubeWidthId             = "cube-width"
	CubeHeightId            = "cube-height"
	CubeDepthId             = "cube-dept"
	BackgroundColorId       = "background-color"
	DirectionalLightColorId = "directional-light-color"
	AmbientLightColorId     = "ambient-light-color"
	RotationSpeedYId        = "rotation-speed-y"
	RotationSpeedXId        = "rotation-speed-x"
)

// Page is a top-level app component.
type Page struct {
	vecty.Core
	Title                 string
	MeshColor             string
	BackgroundColor       string
	DirectionalLightColor string
	AmbientLightColor     string
	MeshWidth             int
	MeshHeight            int
	MeshDepth             int
	RotationSpeedY        int
	RotationSpeedX        int
	SunPosition           [3]float64
	scene                 *three.Scene
	camera                three.PerspectiveCamera
	renderer              *three.WebGLRenderer
	cubeMesh              *three.Mesh
	directionalLight      *three.DirectionalLight
	ambientLight          *three.AmbientLight

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
					p.cubeMesh.Material = three.NewMeshLambertMaterial(params)
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
				case DirectionalLightColorId:
					p.DirectionalLightColor = e.Target.Get("value").String()
					p.directionalLight.Set("color", three.NewColor(p.DirectionalLightColor))
					break
				case AmbientLightColorId:
					p.AmbientLightColor = e.Target.Get("value").String()
					p.ambientLight.Set("color", three.NewColor(p.AmbientLightColor))
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
					p.cubeMesh.Geometry = three.NewBoxGeometry(&three.BoxGeometryParameters{
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
				&components.DisplayButton{
					Id:                 "settings-lock",
					Label:              "Settings",
					TabulationClass:    "main-menu",
					TargetFormSelector: "#form-items-container",
					OffIcon:            "open_in_full",
					OnIcon:             "close_fullscreen",
				},
				elem.Div(
					vecty.Markup(
						prop.ID("form-items-container"),
						vecty.Style("display", "none"),
					),
					&components.DisplayButton{
						Id:                 "cube-display-lock",
						Label:              "Cube Display",
						TabulationClass:    "sub-menu",
						TargetFormSelector: "#cube-display-container",
						OffIcon:            "open_in_full",
						OnIcon:             "close_fullscreen",
					},
					elem.Div(
						vecty.Markup(
							vecty.Class("row"),
							prop.ID("cube-display-container"),
							vecty.Style("display", "none"),
						),
						&components.ColorPicker{Id: CubeColorId, Value: p.MeshColor, Label: "Cube Color:"},
						&components.NumericInput{Id: CubeWidthId, Value: p.MeshWidth, Label: "Cube Width:"},
						&components.NumericInput{Id: CubeHeightId, Value: p.MeshHeight, Label: "Cube Height:"},
						&components.NumericInput{Id: CubeDepthId, Value: p.MeshDepth, Label: "Cube Depth:"},
					),
					&components.DisplayButton{
						Id:                 "cube-rotation-lock",
						Label:              "Cube Rotation",
						TabulationClass:    "sub-menu",
						TargetFormSelector: "#cube-rotation-container",
						OffIcon:            "open_in_full",
						OnIcon:             "close_fullscreen",
					},
					elem.Div(
						vecty.Markup(
							vecty.Class("row"),
							prop.ID("cube-rotation-container"),
							vecty.Style("display", "none"),
						),
						&components.RangeInput{Id: RotationSpeedYId, Value: p.RotationSpeedY, Label: "Y Rotation:", MinValue: -1000, MaxValue: 1000, StepValue: 10},
						&components.RangeInput{Id: RotationSpeedXId, Value: p.RotationSpeedX, Label: "X Rotation:", MinValue: -1000, MaxValue: 1000, StepValue: 10},
					),
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
						elem.Div(
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
								&components.ColorPicker{Id: AmbientLightColorId, Value: p.AmbientLightColor, Label: "Light:"},
							),
						),
						elem.Div(
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
								&components.ColorPicker{Id: DirectionalLightColorId, Value: p.DirectionalLightColor, Label: "Light:"},
							),
						),
					),
					&components.DisplayButton{
						Id:                 "scene-lock",
						Label:              "Scene",
						TabulationClass:    "sub-menu",
						TargetFormSelector: "#scene-container",
						OffIcon:            "open_in_full",
						OnIcon:             "close_fullscreen",
					},
					elem.Div(
						vecty.Markup(
							vecty.Class("row"),
							prop.ID("scene-container"),
							vecty.Style("display", "none"),
						),
						&components.ColorPicker{Id: BackgroundColorId, Value: p.BackgroundColor, Label: "Background:"},
					),
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
	p.renderer.Get("shadowMap").Set("enabled", true)

	// lights
	p.directionalLight = three.NewDirectionalLight(three.NewColor(p.DirectionalLightColor), 1)
	p.directionalLight.Position.Set(p.SunPosition[0], p.SunPosition[1], p.SunPosition[2])
	p.directionalLight.Set("castShadow", true)
	p.directionalLight.Get("shadow").Get("mapSize").Set("width", 1024)
	p.directionalLight.Get("shadow").Get("mapSize").Set("height", 1024)
	p.directionalLight.Get("shadow").Get("camera").Set("left", -300)
	p.directionalLight.Get("shadow").Get("camera").Set("right", 300)
	p.directionalLight.Get("shadow").Get("camera").Set("top", 300)
	p.directionalLight.Get("shadow").Get("camera").Set("bottom", -300)
	p.directionalLight.Get("shadow").Get("camera").Set("far", 1000)
	p.scene.Add(p.directionalLight)

	p.ambientLight = three.NewAmbientLight(three.NewColor(p.AmbientLightColor), 1)
	p.scene.Add(p.ambientLight)

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
	p.cubeMesh = three.NewMesh(geom, mat)
	p.cubeMesh.Set("castShadow", true)
	p.scene.Add(p.cubeMesh)
	p.scene.Background = three.NewColor(p.BackgroundColor)
	// ground texture
	// const RepeatWrapping = 1000;
	textureLoader := three.NewTextureLoader()
	textureLoader.CrossOrigin = "anonymous"
	groundTexture := textureLoader.Load("https://raw.githubusercontent.com/akosgarai/webgl-cube-editor/main/assets/grass.jpg", func(text *js.Object) {
		text.Set("wrapS", 1000)
		text.Set("wrapT", 1000)
		text.Set("anisotropy", 16)
		text.Set("repeat", three.NewVector2(25, 25))
	})
	materialParams := three.NewMaterialParameters()
	materialParams.Map = groundTexture
	groundMaterial := three.NewMeshLambertMaterial(materialParams)
	groundGeom := three.NewBoxGeometry(&three.BoxGeometryParameters{
		Width:  20000,
		Height: 0,
		Depth:  20000,
	})
	groundMesh := three.NewMesh(groundGeom, groundMaterial)
	groundMesh.Position.Set(0, -100, 0)
	groundMesh.Set("receiveShadow", true)
	p.scene.Add(groundMesh)

	// start animation
	p.animate()
}
func (p *Page) shutdown(renderer *three.WebGLRenderer) {
	// After shutdown, we shouldn't use any of these anymore.
	p.scene = nil
	p.camera = three.PerspectiveCamera{}
	p.renderer = nil
	p.cubeMesh = nil
	p.directionalLight = nil
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
	p.cubeMesh.Rotation.Set("y", p.cubeMesh.Rotation.Get("y").Float()+0.0001*float64(p.RotationSpeedY))
	p.cubeMesh.Rotation.Set("x", p.cubeMesh.Rotation.Get("x").Float()+0.0001*float64(p.RotationSpeedX))
	p.renderer.Render(p.scene, p.camera)
}
