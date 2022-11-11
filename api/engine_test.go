package api

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/ravelo-systematic-solutions/fwork/exceptions"
	"github.com/ravelo-systematic-solutions/fwork/response"
	"github.com/ravelo-systematic-solutions/fwork/testutils"
	"log"
	"math/big"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestEngine_Run(t *testing.T) {
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1111),
		Subject: pkix.Name{
			Organization:  []string{"o1"},
			Country:       []string{"c1"},
			Province:      []string{"p1"},
			Locality:      []string{"l1"},
			StreetAddress: []string{"sa1"},
			PostalCode:    []string{"pc1"},
		},
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(0, 0, 1),
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)

	certBytes, _ := x509.CreateCertificate(rand.Reader, cert, cert, &privateKey.PublicKey, privateKey)

	certPEM := new(bytes.Buffer)
	pem.Encode(certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})

	certPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(certPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	serverCert, _ := tls.X509KeyPair(certPEM.Bytes(), certPrivKeyPEM.Bytes())

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
	}

	api := &engine{
		server: http.Server{
			Addr:      ":30000",
			TLSConfig: tlsConfig,
		},
		config: Config{
			Service: Service{
				Id:       "i1",
				Name:     "n1",
				Internal: ":30000",
				External: "https://localhost:30000",
			},
		},
	}

	go func(e *engine) {
		if err := e.server.Shutdown(context.TODO()); err != nil {
			log.Panicf("unable to shutdown: [err: %v]", err)
		}
	}(api)

	//when
	err := api.Run()

	//then
	if err == nil {
		t.Errorf(
			"Run(), error expected",
		)
	}

	ex := err.(*exceptions.Exception)
	if ex.Code != exceptions.ResourceClosedCode {
		t.Errorf(
			"Run(), got %v but want %v",
			ex.Code,
			exceptions.ResourceClosedCode,
		)
	}

	if ex.Message != exceptions.ResourceClosedMessage {
		t.Errorf(
			"Run(), got %v but want %v",
			ex.Message,
			exceptions.ResourceClosedMessage,
		)
	}
}

func TestEngine_GetHandler_success(t *testing.T) {
	//given
	url := "/some-url"
	e := engine{
		routes: make(map[string]Handler, 0),
	}
	key := GenerateEndpointKey(http.MethodGet, url)
	e.routes[key] = func(s Scope) {}

	// when
	handler := e.GetHandler(key)

	//then
	if testutils.GetType(handler) != "github.com/ravelo-systematic-solutions/fwork/api.TestEngine_GetHandler_success.func1" {
		t.Errorf(
			"GetHandler(), got %v but want %v",
			testutils.GetType(handler),
			"fwork/api.TestEngine_GetHandler_success.func1",
		)
	}

}

func TestEngine_GetHandler_NotFounr(t *testing.T) {
	//given
	url := "/some-url"
	e := engine{
		routes: make(map[string]Handler, 0),
	}
	key := GenerateEndpointKey(http.MethodGet, url)
	e.routes[key] = func(s Scope) {}

	// when
	handler := e.GetHandler("invalid-key")

	//then
	if testutils.GetType(handler) != "github.com/ravelo-systematic-solutions/fwork/api.NotFound" {
		t.Errorf(
			"GetHandler(), got %v but want %v",
			testutils.GetType(handler),
			"fwork/api.NotFound",
		)
	}
}

func TestEngine_ServeHTTP_success(t *testing.T) {
	//given
	url := "/some-url"
	e := engine{
		routes: make(map[string]Handler, 0),
	}
	key := GenerateEndpointKey(http.MethodGet, url)
	e.routes[key] = func(s Scope) {
		s.Reply(
			http.StatusAccepted,
			response.Void{},
		)
	}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	expectedResponse := "{}"

	//when
	e.ServeHTTP(w, r)

	//then
	if w.Body.String() != "{}" {
		t.Errorf(
			"ServeHTTP(), got %v but want %v",
			w.Body.String(),
			expectedResponse,
		)
	}
	if w.Code != http.StatusAccepted {
		t.Errorf(
			"ServeHTTP(), got %v but want %v",
			w.Code,
			http.StatusAccepted,
		)
	}
}

func TestEngine_ServeHTTP_Before_success(t *testing.T) {
	//given
	url := "/some-url"
	e := engine{
		routes: make(map[string]Handler, 0),
	}
	key := GenerateEndpointKey(http.MethodGet, url)
	e.routes[key] = func(s Scope) {
		s.Reply(
			http.StatusAccepted,
			response.Void{},
		)
	}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	expectedResponse := "{}"

	//when
	e.ServeHTTP(w, r)

	//then
	if w.Body.String() != "{}" {
		t.Errorf(
			"ServeHTTP(), got %v but want %v",
			w.Body.String(),
			expectedResponse,
		)
	}
	if w.Code != http.StatusAccepted {
		t.Errorf(
			"ServeHTTP(), got %v but want %v",
			w.Code,
			http.StatusAccepted,
		)
	}
}

func TestEngine_ServeHTTP_Before_error(t *testing.T) {}

func TestEngine_ServeHTTP_After_success(t *testing.T) {}

func TestEngine_ServeHTTP_After_error(t *testing.T) {}

func TestNewEngineService(t *testing.T) {
	//given
	certSubject := CertificateSubject{
		Organization:  []string{"o1"},
		Country:       []string{"c1"},
		Province:      []string{"p1"},
		Locality:      []string{"l1"},
		StreetAddress: []string{"sa1"},
		PostalCode:    []string{"pc1"},
		SerialNumber:  123,
		CertNotBefore: time.Now(),
		CertNotAfter:  time.Now().AddDate(0, 0, 1),
	}
	privateKey, _ := GeneratePrivateKey(1024)
	config := Config{
		Service: Service{
			Id:       "i1",
			Name:     "n1",
			Internal: "i1",
			External: "e1",
		},
	}

	// when
	_, err := NewEngine(
		certSubject,
		privateKey,
		config,
	)

	//then
	if err != nil {
		t.Errorf(
			"NewServiceCa(), got error %v",
			err,
		)
	}
}
