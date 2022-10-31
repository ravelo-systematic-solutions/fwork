package api

import (
	"crypto/rsa"
	"fwork/exceptions"
	"log"
	"net/http"
)

//Service holds information
//about a server which references
//an individual conf within
//a cluster of servers
type Service struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Internal string `json:"internal,omitempty"`
	External string `json:"external,omitempty"`
}

type Config struct {
	Service Service
}

type engine struct {
	//server
	server http.Server
	config Config
	routes map[string]Handler

	//interceptors

	//cert
	certSubject CertificateSubject
	privateKey  rsa.PrivateKey
}

// Post is a shortcut for registering a POST Handler
func (e *engine) Post(url string, endpoint Handler) error {
	return e.addEndpoint(http.MethodPost, url, endpoint)
}

// Get is a shortcut for registering a GET Handler
func (e *engine) Get(url string, endpoint Handler) error {
	return e.addEndpoint(http.MethodGet, url, endpoint)
}

// Put is a shortcut for registering a PUT Handler
func (e *engine) Put(url string, endpoint Handler) error {
	return e.addEndpoint(http.MethodPut, url, endpoint)
}

// Delete is a shortcut for registering a DELETE Handler
func (e *engine) Delete(url string, endpoint Handler) error {
	return e.addEndpoint(http.MethodDelete, url, endpoint)
}

// addEndpoint is a shortcut which registers an Api Handler
func (e *engine) addEndpoint(method, url string, endpoint Handler) error {
	key := GenerateEndpointKey(method, url)
	if _, ok := e.routes[key]; ok {
		ex := exceptions.NewBuilder()
		ex.SetCode(exceptions.ResourceDuplicatedCode)
		ex.SetMessage(exceptions.ResourceDuplicatedMessage)
		ex.Include(exceptions.Data{Value: key})

		return ex.Exception()
	}
	e.routes[key] = endpoint
	return nil
}

//ServeHTTP entry point for HTTP requests
func (e *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (e *engine) Run() error {
	log.Printf(
		"Running on %v",
		e.config.Service.External,
	)
	err := e.server.ListenAndServeTLS("", "")

	ex := exceptions.NewBuilder()
	ex.SetCode(exceptions.ResourceClosedCode)
	ex.SetMessage(exceptions.ResourceClosedMessage)
	ex.Include(exceptions.Data{Value: err.Error()})

	return ex.Exception()
}
