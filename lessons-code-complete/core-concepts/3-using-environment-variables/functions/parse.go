
package main

import (
	"os"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type GoogleAppCreds struct {
	Type			string	`json:"type"`
	ProjectID		string	`json:"project_id"`
	PrivateKeyID	string	`json:"private_key_id"`
	PrivateKey		string	`json:"private_key"`
	ClientID		string	`json:"client_id"`
	AuthURI			string	`json:"auth_uri"`
	TokenURI		string	`json:"token_uri"`
	AuthProvider	string	`json:"auth_provider_x509_cert_url"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var gar GoogleAppCreds
	json.Unmarshal([]byte(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")), &gar)
	fmt.Println("projectID: ", gar.ProjectID)
	fmt.Println("tokenURI: ", gar.TokenURI)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Success",
	}, nil
}

func main() {
	lambda.Start(handler)
}