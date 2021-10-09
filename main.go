package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/acm"
)

func main() {
	arn := flag.String("arn", "", "Path to file containing ARN or ARN itself.")
	// outFile := flag.String("out-file", "stdout", "Path to file which will hold certificate.")
	flag.Parse()
	fmt.Println("Getting certificate from ARN: ", *arn)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	acmClient := acm.NewFromConfig(cfg)
	var certInput = &acm.GetCertificateInput{
		CertificateArn: arn,
	}
	getCertificateOutput, err := acmClient.GetCertificate(context.TODO(), certInput)
	if err != nil {
		log.Fatal(err)
	}
	writeOutput(os.Stdout, []byte(*getCertificateOutput.Certificate))
}

func writeOutput(out io.Writer, outBytes []byte) {
	out.Write(outBytes)
}
