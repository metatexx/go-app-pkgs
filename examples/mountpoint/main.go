package main

import (
	"fmt"
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
	Tabs1  []app.UI
	Tabs2  []app.UI
	Active int
	mp     *mountpoint.UI
}

func (c *isolate) OnInit() {
	c.Tabs1 = []app.UI{&tab{Text: "Content A"}, &tab{Text: "Content B"}, &tab{Text: "Content C"}}
	c.Tabs2 = []app.UI{&tab{Text: "Content A"}, &tab{Text: "Content B"}, &tab{Text: "Content C"}}
	c.mp = mountpoint.New(c.Tabs1[0])
}

func (c *isolate) OnNav(ctx app.Context) {
	url := ctx.Page().URL()
	idx, err := strconv.Atoi(url.Fragment)
	if err != nil {
		idx = 0
	}
	c.Active = idx
	c.mp.Switch(c.Tabs1[c.Active])
}

func (c *isolate) Render() app.UI {
	return app.Div().Style("padding", "5px").Body(
		app.Div().Style("padding", "5px").Body(
			app.Span().Text("To see the problem click tabs 0,1,2 and then 0 again."),
			app.Br(),
			app.Span().Text("Notice how the 'intuitive' implementation does not work."),
		),
		app.Div().Body(
			app.A().Style("padding", "5px").Href("#0").Text("Tab 0"),
			app.A().Style("padding", "5px").Href("#1").Text("Tab 1"),
			app.A().Style("padding", "5px").Href("#2").Text("Tab 2"),
			app.Div().Style("margin-top", "10px").Text("Intuitive implementation"),
			app.Div().Style("padding", "5px").Body(c.Tabs2[c.Active]),
			app.Div().Style("margin-top", "10px").Text("Using mountpoint"),
			app.Div().Style("padding", "5px").Body(c.mp.UI()),
		),
	)
}

func main() {
	app.Route("/", &isolate{})
	app.RunWhenOnBrowser()
	http.Handle("/", &app.Handler{
		Name:        "Isolate",
		Description: "Isolated functionality test",
	})
	fmt.Println("open your browser at http://127.0.0.1:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
