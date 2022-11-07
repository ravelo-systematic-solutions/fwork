package api

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/ravelo-systematic-solutions/fwork/exceptions"
	"log"
	"math/big"
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
	i []InterceptorI

	//cert
	certSubject CertificateSubject
	privateKey  rsa.PrivateKey
}

// Controller is a shortcut for registering controllers
func (e *engine) Controller(c Controller) error {

	for k, h := range c.Routes() {
		if _, ok := e.routes[k]; ok {
			ex := exceptions.NewBuilder()
			ex.SetCode(exceptions.ResourceDuplicatedCode)
			ex.SetMessage(exceptions.ResourceDuplicatedMessage)
			ex.Include(exceptions.Data{Value: k})

			return ex.Exception()
		}

		e.routes[k] = h
	}

	return nil
}

//// Post is a shortcut for registering a POST Handler
//func (e *engine) Post(url string, endpoint Handler) error {
//	return e.addEndpoint(http.MethodPost, url, endpoint)
//}
//
//// Get is a shortcut for registering a GET Handler
//func (e *engine) Get(url string, endpoint Handler) error {
//	return e.addEndpoint(http.MethodGet, url, endpoint)
//}
//
//// Put is a shortcut for registering a PUT Handler
//func (e *engine) Put(url string, endpoint Handler) error {
//	return e.addEndpoint(http.MethodPut, url, endpoint)
//}
//
//// Delete is a shortcut for registering a DELETE Handler
//func (e *engine) Delete(url string, endpoint Handler) error {
//	return e.addEndpoint(http.MethodDelete, url, endpoint)
//}
//
//// addEndpoint is a shortcut which registers an Api Handler
//func (e *engine) addEndpoint(method, url string, endpoint Handler) error {
//	key := GenerateEndpointKey(method, url)
//	if _, ok := e.routes[key]; ok {
//		ex := exceptions.NewBuilder()
//		ex.SetCode(exceptions.ResourceDuplicatedCode)
//		ex.SetMessage(exceptions.ResourceDuplicatedMessage)
//		ex.Include(exceptions.Data{Value: key})
//
//		return ex.Exception()
//	}
//	e.routes[key] = endpoint
//	return nil
//}

//ServeHTTP entry point for HTTP requests
func (e *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	errs := make([]exceptions.Exception, 0)

	s := NewScope(w, r)
	key := GenerateEndpointKey(r.Method, r.URL.Path)
	handler := e.GetHandler(key)

	//for _, interceptor := range e.i {
	//	err := interceptor.Before(s)
	//	ex := err.(*exceptions.Exception)
	//	errs = append(errs, *ex)
	//}

	if len(errs) == 0 {
		handler(s)
	}

	//for _, interceptor := range e.i {
	//	err := interceptor.Before(s)
	//	ex := err.(*exceptions.Exception)
	//	errs = append(errs, *ex)
	//}
	//
	//if len(errs) != 0 {
	//	s.JsonRes(http., errs)
	//}

	e.DispatchResponse(s)
}

func (e *engine) DispatchResponse(s *scope) {
	s.w.Header().Set("Access-Control-Allow-Origin", "*")
	s.w.Header().Set("Content-Type", "application/json")
	s.w.WriteHeader(s.s)
	s.w.Write(s.b)
}

//GetHandler retrieves the handler which needs
//to handle the request
func (e *engine) GetHandler(key string) Handler {

	if handler, ok := e.routes[key]; ok {
		return handler
	}

	return NotFound
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

func NewEngine(certSubject CertificateSubject, privateKey *rsa.PrivateKey, config Config) (*engine, error) {

	cert := &x509.Certificate{
		SerialNumber: big.NewInt(certSubject.SerialNumber),
		Subject: pkix.Name{
			Organization:  certSubject.Organization,
			Country:       certSubject.Country,
			Province:      certSubject.Province,
			Locality:      certSubject.Locality,
			StreetAddress: certSubject.StreetAddress,
			PostalCode:    certSubject.PostalCode,
		},
		SubjectKeyId:          []byte(config.Service.Id),
		NotBefore:             certSubject.CertNotBefore,
		NotAfter:              certSubject.CertNotAfter,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: false,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, cert, cert, &privateKey.PublicKey, privateKey)
	if err != nil {
		e := exceptions.NewBuilder()
		e.SetCode(exceptions.ResourceNotGeneratedCode)
		e.SetMessage(exceptions.ResourceNotGeneratedMessage)
		e.Include(exceptions.Data{Value: err.Error()})

		return nil, e.Exception()
	}

	certPEM := new(bytes.Buffer)
	err = pem.Encode(certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	if err != nil {
		e := exceptions.NewBuilder()
		e.SetCode(exceptions.ResourceNotEncodedCode)
		e.SetMessage(exceptions.ResourceNotEncodedMessage)
		e.Include(exceptions.Data{Value: err.Error()})

		return nil, e.Exception()
	}

	certPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(certPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err != nil {
		e := exceptions.NewBuilder()
		e.SetCode(exceptions.ResourceNotEncodedCode)
		e.SetMessage(exceptions.ResourceNotEncodedMessage)
		e.Include(exceptions.Data{Value: err.Error()})

		return nil, e.Exception()
	}

	serverCert, err := tls.X509KeyPair(certPEM.Bytes(), certPrivKeyPEM.Bytes())
	if err != nil {
		e := exceptions.NewBuilder()
		e.SetCode(exceptions.ResourcesNotPairedCode)
		e.SetMessage(exceptions.ResourcesNotPairedMessage)
		e.Include(exceptions.Data{Value: err.Error()})

		return nil, e.Exception()
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
	}

	e := engine{
		server: http.Server{
			Addr:      config.Service.Internal,
			TLSConfig: tlsConfig,
		},
		certSubject: certSubject,
		config:      config,
		routes:      make(map[string]Handler),
	}
	e.server.Handler = &e

	return &e, nil
}
