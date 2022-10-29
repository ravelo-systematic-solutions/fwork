package api

import (
	"testing"
	"time"
)

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
	_, err := NewEngineService(
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
