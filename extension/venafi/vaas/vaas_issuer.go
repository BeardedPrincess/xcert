package vaas

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

func (i *Issuer) SubmitRequest(args interface{}) (string, error) {
	return fmt.Sprintf("Issued by %s: %s", strLongName, args), nil
}

func (i *Issuer) String() string {
	o, _ := json.MarshalIndent(i, "", "  ")
	return string(o)
}

func NewIssuer(conf any) (xcert.CertIssuer, error) {
	v, ok := conf.(map[string]string)
	if !ok {
		return nil, fmt.Errorf("%s: invalid configuration for CertIssuer of type %s", fmt.Sprintf(strFmtComponent, "issuer"), IssuerType)
	}

	return &Issuer{
		AccessToken:  v["accessToken"],
		PolicyFolder: v["policyFolder"],
	}, nil
}
