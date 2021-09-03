package main

import (
	"crypto/x509"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"log"
)

func main() {
	certs, err := x509.SystemCertPool()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Num System Certs: %d\n", len(certs.Subjects()))

	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	stsClient := sts.New(awsSession, aws.NewConfig().WithLogLevel(aws.LogDebugWithSigning))

	response, err := stsClient.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
