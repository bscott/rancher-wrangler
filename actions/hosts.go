package actions

import "github.com/gobuffalo/buffalo"

// HostsShow default implementation.
                func HostsShow(c buffalo.Context) error {
                    return c.Render(200, r.HTML("hosts/show.html"))
                }
                // HostsIndex default implementation.
                func HostsIndex(c buffalo.Context) error {
                    return c.Render(200, r.HTML("hosts/index.html"))
                }
                // HostsCreate default implementation.
                func HostsCreate(c buffalo.Context) error {
                    return c.Render(200, r.HTML("hosts/create.html"))
                }
        