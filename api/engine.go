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
	server      http.Server
	config      Config
	routes      map[string]Handler
	certSubject CertificateSubject
	privateKey  rsa.PrivateKey
}

func (e *engine) Run() error {
	log.Printf(
		"Running on %v",
		e.config.Service.External,
	)
	err := e.server.ListenAndServeTLS("", "")
	if err == http.ErrServerClosed {
		e := exceptions.NewBuilder()
		e.SetCode(exceptions.ResourceClosedCode)
		e.SetMessage(exceptions.ResourceClosedMessage)
		e.Include(exceptions.Data{Value: err.Error()})

		return e.Exception()
	}

	return nil
}