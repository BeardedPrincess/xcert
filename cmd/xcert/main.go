package main

import (
	"fmt"
	"log"

	"github.com/beardedprincess/xcert"
)

func main() {
	conf := &xcert.Config{
		CertIssuer: "tpp",
		Config: map[string]string{
			"policyFolder": "\\VED\\Policy\\Certificates",
			"accessToken":  "332asdfldkfja;ldfj",
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
