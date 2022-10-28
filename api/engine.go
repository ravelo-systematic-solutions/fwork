package api

import (
	"crypto/rsa"
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

func (e *engine) Run() {
	log.Printf(
		"Running on %v",
		e.config.Service.External,
	)
	e.server.ListenAndServeTLS("", "")
}
