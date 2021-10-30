package mountpoint

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type replacer interface {
	app.Composer
	nr() int
	set(app.UI)
}

type mp0 struct {
	app.Compo
	n  int
	ui app.UI
}

func (c *mp0) nr() int {
	return c.n
}

func (c *mp0) set(el app.UI) {
	c.ui = el
}

func (c *mp0) Render() app.UI {
	return c.ui
}

type mp1 struct {
	app.Compo
	n  int
	ui app.UI
}

func (c *mp1) nr() int {
	return c.n
}

func (c *mp1) set(el app.UI) {
	c.ui = el
}

func (c *mp1) Render() app.UI {
	return c.ui
}

//
type UI struct {
	ui replacer
}

func (c *UI) Switch(el app.UI) {
	switch c.ui.nr() {
	case 0:
		c.ui = &mp1{ui: el, n: 1}
	case 1:
		c.ui = &mp0{ui: el}
	}
}

func (m *UI) UI() app.UI {
	return m.ui
}

func New(ui app.UI) *UI {
	return &UI{&mp0{ui: ui}}
}
