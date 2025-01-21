package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Hello World!"),
		app.P().Text("This is hello world sample application."),
		app.Form().Name("user_form").Body(
			app.Label().Text("Name"),
			app.Input().Name("name").Value("Kashif"),
			app.Button().Type("submit").Text("Submit"),
		).OnSubmit(h.onSubmit),
	)
}

func (h *hello) onSubmit(ctx app.Context, e app.Event) {
	fmt.Println("hello onSubmit", e)
	//ctx.ShowAppInstallPrompt()
}
