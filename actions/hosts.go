package actions

import "github.com/gobuffalo/buffalo"

type HostsResource struct {
	buffalo.Resource
}

func init() {
	var resource buffalo.Resource
	resource = &HostsResource{&buffalo.BaseResource{}}
	App().Resource("/hosts", resource)
}

// List default implementation.
func (v *HostsResource) List(c buffalo.Context) error {
	return c.Render(200, r.String("Hosts#List"))
}

// Show default implementation.
func (v *HostsResource) Show(c buffalo.Context) error {
	return c.Render(200, r.String("Hosts#Show"))
}

// New default implementation.
func (v *HostsResource) New(c buffalo.Context) error {
	return c.Render(200, r.String("Hosts#New"))
}

// Create default implementation.
func (v *HostsResource) Create(c buffalo.Context) error {
	return c.Render(200, r.String("Hosts#Create"))
}

// Edit default implementation.
func (v *HostsResource) Edit(c buffalo.Context) error {
	return c.Render(200, r.String("Hosts#Edit"))
}

// Update default implementation.
func (v *HostsResource) Update(c buffalo.Context) error {
	return c.Render(200, r.String("Hosts#Update"))
}

// Destroy default implementation.
func (v *HostsResource) Destroy(c buffalo.Context) error {
	return c.Render(200, r.String("Hosts#Destroy"))
}
