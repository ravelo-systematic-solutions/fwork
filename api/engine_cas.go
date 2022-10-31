package api

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fwork/exceptions"
	"math/big"
	"net/http"
)

func NewServiceCa(certSubject CertificateSubject, privateKey *rsa.PrivateKey, config Config) (*engine, error) {

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
		NotBefore:             certSubject.CertNotBefore,
		NotAfter:              certSubject.CertNotAfter,
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
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

	return &engine{
		config: config,
		server: http.Server{
			Addr:      config.Service.Internal,
			TLSConfig: tlsConfig,
		},
		i:           make([]InterceptorI, 0),
		certSubject: certSubject,
	}, nil
}
