package api

import (
	"time"
)

type CertificateSubject struct {
	Organization  []string
	Country       []string
	Province      []string
	Locality      []string
	StreetAddress []string
	PostalCode    []string
	SerialNumber  int64
	CertNotBefore time.Time
	CertNotAfter  time.Time
}
