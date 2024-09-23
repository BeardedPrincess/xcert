package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/beardedprincess/xcert"
	_ "github.com/beardedprincess/xcert/extension/venafi/tpp"
	_ "github.com/beardedprincess/xcert/extension/venafi/vaas"
)

var (
	versionBuildTimeStamp string
	versionString         string
	commit                string
)

func main() {
	// Propogate versions strings to xcert base package to be shared globally
	xcert.VersionString = versionString
	xcert.VersionBuildTimeStamp = versionBuildTimeStamp
	xcert.Commit = commit

	// xcert.SetLogger(log.Printf, log.Printf)

	fmt.Printf("Version: %v\n\n", xcert.GetFormattedVersionString())

	vi := strings.Join(xcert.ValidIssuers(), ", ")
	log.Printf("Valid issuers: %v\n", vi)

	log.Printf("%s is valid?: %v", "foo", xcert.IsValidIssuer("foo"))

	conf := &xcert.Config{
		CertIssuer: "tpp",
		Config: map[string]any{
			"policyFolder": "\\VED\\Policy\\Certificates",
			"accessToken":  "afadfadsfadsfa",
		},
	}

	client, err := conf.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	r, err := client.Issuer.SubmitRequest("3dsfasdfaf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(client.Issuer)
	fmt.Println(r)

}
