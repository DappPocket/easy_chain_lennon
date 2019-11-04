package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	isDev := envy.Get("IS_DEV", "false")
	if isDev == "true" {
		return c.Render(200, r.HTML("index.html"))
	} else if isDev ==  "prod" {
		return c.Render(200, r.HTML("indexpd.html", "frontend.html"))
	} else {
		return c.Render(200, r.JSON(map[string]string{
			"message": "Hello $__$. Welcome to easy ChainLife!",
			"donate_me": "0x99993AE2576B71AA30fdB9dA879765620AbcB7F3",
		}))
	}
}
