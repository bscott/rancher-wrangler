package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/bscott/rancher-wrangler/models"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	var hosts models.Hosts
	models.DB.All(&hosts)


	c.Set("hosts", hosts)
	return c.Render(200, r.HTML("index.html"))
}


