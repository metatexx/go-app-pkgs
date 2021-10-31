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

type UI struct {
	ui replacer
}

// Switch lets you switch the app.UI component with another one.
// It guarantees that the former component gets dismounted and
// the new one gets mounted in place of the old one.
func (c *UI) Switch(el app.UI) {
	if c.ui == el {
		return
	}
	if el.Mounted() {
		return
	}
	switch c.ui.nr() {
	case 0:
		c.ui = &mp1{ui: el, n: 1}
	case 1:
		c.ui = &mp0{ui: el}
	}
}

// UI returns the reference to the current mounted app.UI
func (m *UI) UI() app.UI {
	return m.ui
}

// New creates a new mountpoint for switching app.UI components
func New(ui app.UI) *UI {
	return &UI{&mp0{ui: ui}}
}
