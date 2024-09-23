// Copyright (c) BeardedPrincess 2024
// SPDX-License-Identifier: MPL-2.0

package xcert

import (
	"fmt"
	"strings"
)

type CertIssuerType string

// Responsible for issuing X509 certificates and implements the CertIssuer interface
type CertIssuer interface {
	GetType() CertIssuerType
	GetName() string
	SubmitRequest(args interface{}) (string, error)
	String() string
}

type CertIssuerFactory func(conf any) (CertIssuer, error)

var certIssuers = make(map[CertIssuerType]CertIssuerFactory)

func ValidIssuers() []string {
	var ret []string
	for is := range certIssuers {
		ret = append(ret, string(is))
	}
	return ret
}

func IsValidIssuer(t CertIssuerType) bool {
	_, ok := certIssuers[t]
	return ok
}

func Register(t CertIssuerType, f CertIssuerFactory) error {
	const op = "/certIssuer/Register"
	if f == nil {
		return &Error{
			Op:  fmt.Sprintf(strFmtComponent, op),
			Err: fmt.Errorf("CertIssuer %s failed registration. An invalid factory function was provided", t),
		}
	}

	_, registered := certIssuers[t]
	if registered {
		LogWarn("%s is already a registered issuer", string(t))
		return nil
	}

	certIssuers[t] = f
	LogDebug("registered issuer: %s", t)
	return nil
}

func NewCertIssuer(t CertIssuerType, conf any) (CertIssuer, error) {
	const op = "/certIssuer/NewCertIssuer"
	if string(t) == "" {
		return nil, &Error{
			Op:  fmt.Sprintf(strFmtComponent, op),
			Err: fmt.Errorf("issuerType cannot be empty"),
		}
	}

	certIssuerFactory, ok := certIssuers[t]
	if !ok {
		// Specified cert issuer is not registered
		availableIssuers := make([]string, 0)
		for k := range certIssuers {
			availableIssuers = append(availableIssuers, string(k))
		}
		return nil, fmt.Errorf("invalid issuer type. Must be one of: \"%s\"", strings.Join(availableIssuers, ", "))
	}

	return certIssuerFactory(conf)
}
