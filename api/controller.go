package api

import "net/http"

type ControllerI interface {
	ResourceName() string
}

type Handlers struct {
	Get    Handler
	Post   Handler
	Put    Handler
	Patch  Handler
	Delete Handler
}

type Controller struct {
	resource string
	routes   map[string]Handler
}

func (c *Controller) ResourceName() string {
	return c.resource
}

func (c *Controller) GetHandler(method, url string) Handler {
	key := GenerateEndpointKey(http.MethodGet, c.resource)

	if h, ok := c.routes[key]; ok {
		return h
	}

	return NotFound
}

func NewController(url string, handlers Handlers) Controller {
	c := Controller{
		resource: url,
		routes:   make(map[string]Handler),
	}

	if handler := handlers.Get; handler != nil {
		key := GenerateEndpointKey(http.MethodGet, c.resource)
		c.routes[key] = handler
	}

	if handler := handlers.Post; handler != nil {
		key := GenerateEndpointKey(http.MethodPost, c.resource)
		c.routes[key] = handler
	}

	if handler := handlers.Put; handler != nil {
		key := GenerateEndpointKey(http.MethodPut, c.resource)
		c.routes[key] = handler
	}

	return c
}
