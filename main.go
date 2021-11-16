package main

import (
	"fmt"
	"log"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	"github.com/watson-developer-cloud/go-sdk/v2/assistantv2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env File")
	}
	ver := os.Getenv("version")
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("apiKey"),
	}

	options := &assistantv2.AssistantV2Options{
		Version:       &ver,
		Authenticator: authenticator,
	}

	assistant, assistantErr := assistantv2.NewAssistantV2(options)

	if assistantErr != nil {
		panic(assistantErr)
	}

	assistant.SetServiceURL(os.Getenv("serviceUrl"))

	result, _, responseErr := assistant.CreateSession(assistant.
		NewCreateSessionOptions(os.Getenv("assistantId")))
	if responseErr != nil {
		panic(responseErr)
	}
	fmt.Println(*result.SessionID)
}
