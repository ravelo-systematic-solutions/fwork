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
	"log"
	"math/big"
	"net"
	"net/http"
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
	if err != nil {
		t.Errorf(
			"Run(), unexpected error: %v",
			err,
		)
	}
}
