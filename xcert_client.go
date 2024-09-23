// Copyright (c) BeardedPrincess 2024
// SPDX-License-Identifier: MPL-2.0

package xcert

type Client struct {
	Issuer CertIssuer
}

func (c *Config) NewClient() (*Client, error) {
	const op = "Config.NewClient"
	issuer, err := NewCertIssuer(CertIssuerType(c.CertIssuer), c.Config)
	if err != nil {
		return nil, &Error{
			Op:  op,
			Err: err,
		}
	}

	return &Client{Issuer: issuer}, nil
}
