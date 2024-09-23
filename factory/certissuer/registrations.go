package factorycertissuer

import (
	"github.com/beardedprincess/xcert/extension/venafi/tpp"
	"github.com/beardedprincess/xcert/shared/certissuer"
)

func init() {
	Register(certissuer.CertIssuerType(tpp.IssuerType), CertIssuerFactory(tpp.NewIssuer))
}
