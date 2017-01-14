package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/bscott/rancher-wrangler/models"
)

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
		// HostsNew implements creation of a new host.
		func HostsNew (c buffalo.Context) error {
			name := c.Param("name")
			desc := c.Param("description")
			url := c.Param("url")
			api_version := c.Param("api_version")
			access_key := c.Param("access_key")
			secret_key := c.Param("secret_key")
			host := models.Host{Name: name,Description: desc,Url: url,ApiVersion:api_version,AccessKey:access_key,SecretKey:secret_key}
			models.DB.Create(&host)
			c.LogField("Host Created:", host)
			return c.Render(200, r.String("Host Created"))

		}
        