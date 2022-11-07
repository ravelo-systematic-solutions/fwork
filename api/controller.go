package api

import "net/http"

type Controller interface {
	Url() string
	Routes() map[string]Handler
	GetHandler(method, url string) Handler
}

type Endpoints struct {
	Get    Handler
	Post   Handler
	Put    Handler
	Patch  Handler
	Delete Handler
}

type Resource struct {
	url    string
	routes map[string]Handler
}

func (r *Resource) Url() string {
	return r.url
}

func (r *Resource) Routes() map[string]Handler {
	return r.routes
}

func (r *Resource) GetHandler(method, url string) Handler {
	key := GenerateEndpointKey(method, url)

	if h, ok := r.routes[key]; ok {
		return h
	}

	return NotFound
}

func NewResource(url string, handlers Endpoints) Resource {
	c := Resource{
		url:    url,
		routes: make(map[string]Handler),
	}

	if handler := handlers.Get; handler != nil {
		key := GenerateEndpointKey(http.MethodGet, c.url)
		c.routes[key] = handler
	}

	if handler := handlers.Post; handler != nil {
		key := GenerateEndpointKey(http.MethodPost, c.url)
		c.routes[key] = handler
	}

	if handler := handlers.Put; handler != nil {
		key := GenerateEndpointKey(http.MethodPut, c.url)
		c.routes[key] = handler
	}

	if handler := handlers.Patch; handler != nil {
		key := GenerateEndpointKey(http.MethodPatch, c.url)
		c.routes[key] = handler
	}

	if handler := handlers.Delete; handler != nil {
		key := GenerateEndpointKey(http.MethodDelete, c.url)
		c.routes[key] = handler
	}

	return c
}
