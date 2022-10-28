package main

import (
	"fwork/api"
	"log"
	"time"
)

func main() {

	subject := api.CertificateSubject{
		SerialNumber:  int64(time.Now().Year()),
		Organization:  []string{"Ravelo Systematic Solution Inc."},
		Country:       []string{"CA"},
		Province:      []string{"BC"},
		Locality:      []string{"Vancouver"},
		StreetAddress: []string{"2289 E 1st Av"},
		PostalCode:    []string{"V5M 0G2"},
		CertNotBefore: time.Now(),
		CertNotAfter:  time.Now().AddDate(10, 0, 0),
	}
	privateKey, err := api.GeneratePrivateKey(4096)

	if err != nil {
		log.Printf("failed to generate private key: %v", err)
		return
	}

	server, err := api.NewServiceCa(subject, privateKey)

	if err != nil {
		log.Printf("failed to instantiate service: %v", err)
		return
	}

	server.Run()
}
