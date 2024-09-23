package factorycertissuer

import (
	"fmt"
	"log"
	"strings"

	"github.com/beardedprincess/xcert/shared/certissuer"
)

var certIssuers = make(map[certissuer.CertIssuerType]CertIssuerFactory)

type CertIssuerFactory func(conf any) (certissuer.CertIssuer, error)

func Register(t certissuer.CertIssuerType, factory CertIssuerFactory) {
	if factory == nil {
		log.Panicf("CertIssuer %s failed registration. An invalid factory function was provided", t)
	}

	_, registered := certIssuers[t]
	if registered {
		log.Printf("CertIssuer %s is already registered. Ignoring.", t)
		return
	}

	certIssuers[t] = factory
	log.Default().Printf("Registered CertIssuer: %s", t)
}

func NewCertIssuer(t certissuer.CertIssuerType, conf interface{}) (certissuer.CertIssuer, error) {
	if string(t) == "" {
		t = "tpp"
	}

	certIssuerFactory, ok := certIssuers[t]
	if !ok {
		// Specified cert issuer is not registered
		availableIssuers := make([]string, 0)
		for k := range certIssuers {
			availableIssuers = append(availableIssuers, string(k))
		}
		return nil, fmt.Errorf(fmt.Sprintf("Invalid issuer type. Must be one of: \"%s\"", strings.Join(availableIssuers, ", ")))
	}

	return certIssuerFactory(conf)
}
