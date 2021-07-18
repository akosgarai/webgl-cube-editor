package main

import (
	"strconv"
	"strings"

	"github.com/akosgarai/webgl-cube-editor/pkg/components"
	"github.com/akosgarai/webgl-cube-editor/pkg/forms"
	"github.com/akosgarai/webgl-cube-editor/pkg/wglrenderer"
	"github.com/gmlewis/go-threejs/three"
	"github.com/gopherjs/gopherjs/js"
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
)

const (
	CubeColorId                 = "cube-color"
	CubeWidthId                 = "cube-width"
	CubeHeightId                = "cube-height"
	CubeDepthId                 = "cube-dept"
	BackgroundColorId           = "background-color"
	DirectionalLightColorId     = "directional-light-color"
	DirectionalLightIntensityId = "directional-light-intensity"
	AmbientLightColorId         = "ambient-light-color"
	AmbientLightIntensityId     = "ambient-light-intensity"
	RotationSpeedYId            = "rotation-speed-y"
	RotationSpeedXId            = "rotation-speed-x"
)

var (
	threejs = three.New()
)

// Page is a top-level app component.
type Page struct {
	vecty.Core

	Title                     string
	MeshColor                 string
	BackgroundColor           string
	DirectionalLightColor     string
	DirectionalLightIntensity float64
	AmbientLightColor         string
	AmbientLightIntensity     float64
	MeshWidth                 int
	MeshHeight                int
	MeshDepth                 int
	RotationSpeedY            int
	RotationSpeedX            int
	SunPosition               [3]float64

	scene            *three.Scene
	camera           *three.PerspectiveCamera
	renderer         *three.WebGLRenderer
	cubeMesh         *three.Mesh
	directionalLight *three.DirectionalLight
	ambientLight     *three.AmbientLight

	canvasWidth  float64
	canvasHeight float64
}

func ColorpicerValueToInt(color string) int {
	col, _ := strconv.ParseInt(strings.Trim(color, "#"), 16, 32)
	return int(col)
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
					p.cubeMesh.JSObject().Get("material").Set("color", threejs.NewColor(ColorpicerValueToInt(p.MeshColor)))
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
					p.scene.JSObject().Set("background", threejs.NewColor(ColorpicerValueToInt(p.BackgroundColor)))
					break
				case DirectionalLightColorId:
					p.DirectionalLightColor = e.Target.Get("value").String()
					p.directionalLight.JSObject().Set("color", threejs.NewColor(ColorpicerValueToInt(p.DirectionalLightColor)))
					break
				case DirectionalLightIntensityId:
					p.DirectionalLightIntensity, _ = strconv.ParseFloat(e.Target.Get("value").String(), 64)
					p.directionalLight.JSObject().Set("intensity", p.DirectionalLightIntensity)
					break
				case AmbientLightColorId:
					p.AmbientLightColor = e.Target.Get("value").String()
					p.ambientLight.JSObject().Set("color", threejs.NewColor(ColorpicerValueToInt(p.AmbientLightColor)))
					break
				case AmbientLightIntensityId:
					p.AmbientLightIntensity, _ = strconv.ParseFloat(e.Target.Get("value").String(), 64)
					p.ambientLight.JSObject().Set("intensity", p.AmbientLightIntensity)
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
					p.cubeMesh.JSObject().Set("geometry", threejs.NewBoxGeometry(float64(p.MeshWidth), float64(p.MeshHeight), float64(p.MeshDepth), nil))
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
					&forms.CubeDisplay{
						CubeColorId:  CubeColorId,
						CubeWidthId:  CubeWidthId,
						CubeHeightId: CubeHeightId,
						CubeDepthId:  CubeDepthId,
						CubeColor:    p.MeshColor,
						CubeWidth:    p.MeshWidth,
						CubeHeight:   p.MeshHeight,
						CubeDepth:    p.MeshDepth,
					},
					&forms.CubeRotation{
						RotationComponentXId: RotationSpeedXId,
						RotationXValue:       p.RotationSpeedX,
						RotationComponentYId: RotationSpeedYId,
						RotationYValue:       p.RotationSpeedY,
					},
					&forms.Lightsource{
						AmbientLightColorId:         AmbientLightColorId,
						AmbientLightColor:           p.AmbientLightColor,
						AmbientLightIntensityId:     AmbientLightIntensityId,
						AmbientLightIntensity:       p.AmbientLightIntensity,
						DirectionalLightColorId:     DirectionalLightColorId,
						DirectionalLightColor:       p.DirectionalLightColor,
						DirectionalLightIntensity:   p.DirectionalLightIntensity,
						DirectionalLightIntensityId: DirectionalLightIntensityId,
					},
					&forms.Scene{
						BackgroundColorId: BackgroundColorId,
						BackgroundColor:   p.BackgroundColor,
					},
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
	p.camera = threejs.NewPerspectiveCamera(70, p.canvasWidth/p.canvasHeight, 1, 1000)
	p.camera.Position().Set(0, 0, 400)
	p.scene = threejs.NewScene()

	p.renderer = p.renderer.SetPixelRatio(devicePixelRatio)
	p.renderer.SetSize(p.canvasWidth, p.canvasHeight, true)
	p.renderer.JSObject().Get("shadowMap").Set("enabled", true)
	// lights
	p.directionalLight = threejs.NewDirectionalLight(float64(ColorpicerValueToInt(p.DirectionalLightColor)), p.DirectionalLightIntensity)
	p.directionalLight.Position().Set(p.SunPosition[0], p.SunPosition[1], p.SunPosition[2])
	p.directionalLight.JSObject().Set("castShadow", true)
	p.directionalLight.JSObject().Get("shadow").Get("mapSize").Set("width", 1024)
	p.directionalLight.JSObject().Get("shadow").Get("mapSize").Set("height", 1024)
	p.directionalLight.JSObject().Get("shadow").Get("camera").Set("left", -300)
	p.directionalLight.JSObject().Get("shadow").Get("camera").Set("right", 300)
	p.directionalLight.JSObject().Get("shadow").Get("camera").Set("top", 300)
	p.directionalLight.JSObject().Get("shadow").Get("camera").Set("bottom", -300)
	p.directionalLight.JSObject().Get("shadow").Get("camera").Set("far", 1000)
	p.scene.Add(p.directionalLight)
	p.ambientLight = threejs.NewAmbientLight(ColorpicerValueToInt(p.AmbientLightColor), p.AmbientLightIntensity)
	p.scene.Add(p.ambientLight)

	// material
	mat := threejs.NewMeshLambertMaterial(three.MeshLambertMaterialOpts{"color": threejs.NewColor(ColorpicerValueToInt(p.MeshColor))})

	// cube object
	geom := threejs.NewBoxGeometry(float64(p.MeshWidth), float64(p.MeshHeight), float64(p.MeshDepth), nil)
	p.cubeMesh = threejs.NewMesh(geom, mat)
	p.cubeMesh.JSObject().Set("castShadow", true)
	p.scene.Add(p.cubeMesh)
	p.scene.JSObject().Set("background", threejs.NewColor(ColorpicerValueToInt(p.BackgroundColor)))
	// ground texture
	// const RepeatWrapping = 1000;
	textureLoader := threejs.NewTextureLoader()
	textureLoader.SetCrossOrigin("anonymous")
	groundTexture := textureLoader.Load("https://raw.githubusercontent.com/akosgarai/webgl-cube-editor/main/assets/grass.jpg", nil, nil, nil)
	groundTexture.SetWrapS(1000)
	groundTexture.SetWrapT(1000)
	groundTexture.SetAnisotropy(16)
	groundTexture.JSObject().Get("repeat").Set("x", 25)
	groundTexture.JSObject().Get("repeat").Set("y", 25)
	groundMesh := threejs.NewMesh(threejs.NewBoxGeometry(20000, 0, 20000, nil), threejs.NewMeshLambertMaterial(three.MeshLambertMaterialOpts{"map": groundTexture.JSObject()}))
	groundMesh.Position().Set(0, -100, 0)
	groundMesh.JSObject().Set("receiveShadow", true)
	p.scene.Add(groundMesh)
	// start animation
	p.animate()
}
func (p *Page) shutdown(renderer *three.WebGLRenderer) {
	// After shutdown, we shouldn't use any of these anymore.
	p.scene = nil
	p.camera = &three.PerspectiveCamera{}
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
		p.camera.JSObject().Set("aspect", windowWidth/p.canvasHeight)
		p.camera.UpdateProjectionMatrix()
		p.renderer.SetSize(windowWidth, p.canvasHeight, false)
		p.canvasWidth = windowWidth
	}
	js.Global.Call("requestAnimationFrame", p.animate)
	p.cubeMesh.JSObject().Get("rotation").Set("y", p.cubeMesh.JSObject().Get("rotation").Get("y").Float()+0.0001*float64(p.RotationSpeedY))
	p.cubeMesh.JSObject().Get("rotation").Set("x", p.cubeMesh.JSObject().Get("rotation").Get("x").Float()+0.0001*float64(p.RotationSpeedX))

	p.renderer.Render(p.scene, p.camera, nil)
}
