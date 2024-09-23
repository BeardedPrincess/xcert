package tpp

import (
	"encoding/json"
	"fmt"

	certissuer "github.com/beardedprincess/xcert/shared/certissuer"
)

var issuerType certissuer.CertIssuerType = "tpp"

var IssuerType certissuer.CertIssuerType = issuerType

type Issuer struct {
	AccessToken  string `json:"accessToken"`
	PolicyFolder string `json:"policyFolder"`
}

func init() {

}

func (i *Issuer) GetType() certissuer.CertIssuerType {
	return issuerType
}

func (i *Issuer) GetName() string {
	return string(issuerType)
}

func (i *Issuer) SubmitRequest(args interface{}) (string, error) {
	return fmt.Sprintf("Issued by TPP: %s", args), nil
}

func (i *Issuer) String() string {
	o, _ := json.MarshalIndent(i, "", "  ")
	return string(o)
}

func NewIssuer(conf any) (certissuer.CertIssuer, error) {
	v, ok := conf.(map[string]string)
	if !ok {
		return nil, fmt.Errorf("invalid configuration for CertIssuer of type %s", issuerType)
	}

	return &Issuer{
		AccessToken:  v["accessToken"],
		PolicyFolder: v["policyFolder"],
	}, nil
}
