package tpp

import (
	"encoding/json"
	"fmt"

	"github.com/beardedprincess/xcert"
)

// Ensure this Issuer complies with the xcert.CertIssuer Interface
var _ xcert.CertIssuer = &Issuer{}

var IssuerType xcert.CertIssuerType = strType

type Issuer struct {
	AccessToken  string `json:"accessToken"`
	PolicyFolder string `json:"policyFolder"`
}

func init() {
	xcert.Register(IssuerType, xcert.CertIssuerFactory(NewIssuer))
}

func (i *Issuer) GetType() xcert.CertIssuerType {
	return IssuerType
}

func (i *Issuer) GetName() string {
	return string(IssuerType)
}

func (i *Issuer) SubmitRequest(r any) (string, error) {
	const op = "tpp.Issuer.submitRequest"
	_ = op
	xcert.LogDebug("[%s] processing request: %v", op, r)
	return fmt.Sprintf("Issued by %s: %s", strLongName, r), nil
}

func (i *Issuer) String() string {
	o, _ := json.MarshalIndent(i, "", "  ")
	return string(o)
}

func NewIssuer(conf any) (xcert.CertIssuer, error) {
	const op = "tpp.Issuer.NewIssuer"
	v, ok := conf.(map[string]any)
	if !ok {
		return nil, &xcert.Error{
			Op:      op,
			Message: EINVALIDCONFIG,
			Code:    xcert.EINVALID,
		}
	}

	t, ok := v["accessToken"].(string)
	if !ok {
		return nil, &xcert.Error{
			Op:      op,
			Message: EACCESSTOKENREQD,
			Code:    xcert.EINVALID,
		}
	}

	p, ok := v["policyFolder"].(string)
	if !ok {
		return nil, &xcert.Error{
			Op:      op,
			Message: EPOLICYFOLDERREQD,
			Code:    xcert.EINVALID,
		}
	}

	return &Issuer{
		AccessToken:  t,
		PolicyFolder: p,
	}, nil
}
