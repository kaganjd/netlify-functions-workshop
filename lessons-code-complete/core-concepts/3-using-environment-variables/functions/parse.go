package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type GoogleAppCreds struct {
	Type                string `json:"type"`
	ProjectID           string `json:"project_id"`
	PrivateKeyID        string `json:"private_key_id"`
	PrivateKey          string `json:"private_key"`
	ClientEmail         string `json:"client_email"`
	ClientID            string `json:"client_id"`
	AuthURI             string `json:"auth_uri"`
	TokenURI            string `json:"token_uri"`
	AuthProviderCertURL string `json:"auth_provider_x509_cert_url"`
	ClientCertURL       string `json:"client_x509_cert_url"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var gar GoogleAppCreds

	// Convert string to JSON and log a few values
	json.Unmarshal([]byte(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")), &gar)
	fmt.Println("projectID: ", gar.ProjectID)
	fmt.Println("privateKey: ", gar.PrivateKey)
	fmt.Println("tokenURI: ", gar.TokenURI)

	// Convert JSON back to string to return as HTTP body
	b, err := json.Marshal(gar)
	if err != nil {
		panic(err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(b),
	}, nil
}

func main() {
	lambda.Start(handler)
}
