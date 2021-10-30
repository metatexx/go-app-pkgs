package main

import (
	"github.com/metatexx/go-app-pkgs/mountpoint"
	"log"
	"net/http"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type tab struct {
	app.Compo
	Text string
}

func (c *tab) Render() app.UI {
	return app.Div().Body(
		app.Div().Text(c.Text),
	)
}

type isolate struct {
	app.Compo
	Tabs    []app.UI
	Active  int
	mp *mountpoint.UI
}

var _ app.Mounter = (*isolate)(nil)

func (c *isolate) OnMount(ctx app.Context) {
	c.Tabs = []app.UI{&tab{Text: "Content A"}, &tab{Text: "Content B"}, &tab{Text: "Content C"}}
	c.mp = mountpoint.New(c.Tabs[0])
}

func (c *isolate) OnNav(ctx app.Context) {
	url := ctx.Page().URL()
	idx, err := strconv.Atoi(url.Fragment)
	if err != nil {
		idx = 0
	}
	c.Active = idx
	c.mp.Switch(c.Tabs[c.Active])
}

func (c *isolate) Render() app.UI {
	if c.Tabs == nil {
		return app.Div().Text("Unmounted")
	}
	return app.Div().Body(
		app.A().Style("padding", "5px").Href("#0").Text("Tab 0"),
		app.A().Style("padding", "5px").Href("#1").Text("Tab 1"),
		app.A().Style("padding", "5px").Href("#2").Text("Tab 2"),
		app.Div().Style("padding", "5px").Body(c.mp.UI()),
	)
}

func main() {
	app.Route("/", &isolate{})
	app.RunWhenOnBrowser()
	http.Handle("/", &app.Handler{
		Name:        "Isolate",
		Description: "Isolated functionality test",
	})
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}