package xcert

import (
	factorycertissuer "github.com/beardedprincess/xcert/factory/certissuer"
	"github.com/beardedprincess/xcert/shared/certissuer"
)

type Client struct {
	Issuer certissuer.CertIssuer
}

func (c *Config) NewClient() (*Client, error) {

	issuer, err := factorycertissuer.NewCertIssuer(certissuer.CertIssuerType(c.CertIssuer), c.Config)
	if err != nil {
		return nil, err
	}

	return &Client{Issuer: issuer}, nil
}
