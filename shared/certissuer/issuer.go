package certissuer

type CertIssuerType string

type CertIssuer interface {
	GetType() CertIssuerType
	GetName() string
	SubmitRequest(args interface{}) (string, error)
	String() string
}
