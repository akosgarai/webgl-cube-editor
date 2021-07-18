package wglrenderer

import (
	"github.com/gmlewis/go-threejs/three"
	"github.com/gopherjs/gopherjs/js"
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type webGLRenderer struct {
	vecty.Core
	opts   WebGLOptions          `vecty:"prop"`
	markup []vecty.MarkupOrChild `vecty:"prop"`

	canvas   *vecty.HTML
	renderer *three.WebGLRenderer
}

// Mount implements the vecty.Mounter interface.
func (r *webGLRenderer) Mount() {
	r.renderer = newWebGLRenderer(&webGLRendererParameters{Canvas: r.canvas.Node()})
	r.opts.Init(r.renderer)
}

// Unmount implements the vecty.Unmounter interface.
func (r *webGLRenderer) Unmount() {
	if r.opts.Shutdown != nil {
		r.opts.Shutdown(r.renderer)
	}
}

// Render implements the vecty.Component interface.
func (r *webGLRenderer) Render() vecty.ComponentOrHTML {
	r.canvas = elem.Canvas(r.markup...)
	return r.canvas
}

// WebGLOptions represent options for the WebGLRenderer component.
type WebGLOptions struct {
	// Init is called when the three.js WebGLRenderer has been created.
	//
	// This can happen multiple times during the lifecycle of an application
	// if the Vecty WebGLRenderer component was unmounted and mounted again,
	// e.g. due to navigating to a different page and back again.
	Init func(r *three.WebGLRenderer)

	// Shutdown is called before the canvas associated with the three.js
	// WebGLRenderer will be destroyed. For example, when your Vecty
	// application no longer renders the WebGLRenderer component and it is
	// being unmounted.
	Shutdown func(r *three.WebGLRenderer)

	// TODO(slimsag): allow specifying other parameters like context, precision,
	// etc. from three.js WebGLRenderer constructor here:
	// https://threejs.org/docs/#api/renderers/WebGLRenderer
}

// WebGLRenderer returns a Vecty component that initializes a three.js WebGL renderer for
// use in a Vecty application.
func WebGLRenderer(opts WebGLOptions, markup ...vecty.MarkupOrChild) vecty.Component {
	if opts.Init == nil {
		panic("vthree: Renderer: must specify opts.Init")
	}
	return &webGLRenderer{
		opts:   opts,
		markup: markup,
	}
}

type webGLRendererParameters struct {
	Canvas interface{}
}

// Note: We can't use three.NewWebGLRenderer because it doesn't allow
// specifying any parameters yet. Easy enough to just call ourself, though.
func newWebGLRenderer(parameters *webGLRendererParameters) *three.WebGLRenderer {
	//return three.WebGLRendererFromJSObject(js.MakeWrapper(map[string]interface{}{"canvas": parameters.Canvas}))
	return three.WebGLRendererFromJSObject(js.Global.Get("THREE").Get("WebGLRenderer").New(map[string]interface{}{
		"canvas": parameters.Canvas,
	}))
}
