/*package main

import (
	"encoding/json"
	"fmt"
















	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/v2/assistantv2"
)

func main() {
	authenticator := &core.IamAuthenticator{
		ApiKey: "HXoxv39NR1RyhzCxsJYhNCm9tFKbL1kni0mmnpJW7Fn5",
	}

	options := &assistantv2.AssistantV2Options{
		Version:       "2021-06-14",
		Authenticator: authenticator,
	}

	assistant, assistantErr := assistantv2.NewAssistantV2(options)

	if assistantErr != nil {
		panic(assistantErr)
	}

	assistant.SetServiceURL("https://api.eu-gb.assistant.watson.cloud.ibm.com/instances/c630716a-83de-4e23-bb90-d51e23c6c5a0/v2/assistants/1e108d47-e09f-4b70-b0c4-a8b1e0ac1813/sessions")

	result, _, responseErr := assistant.Message(
		&assistantv2.MessageOptions{
			AssistantID: core.StringPtr("1e108d47-e09f-4b70-b0c4-a8b1e0ac1813"),
			SessionID:   core.StringPtr("{session_id}"),
			Input: &assistantv2.MessageInput{
				MessageType: core.StringPtr("text"),
				Text:        core.StringPtr("Hello"),
			},
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}
	b, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(b))
}
*/
package main

import (
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/v2/assistantv2"
)

func main() {
	ver := "2021-11-16"
	authenticator := &core.IamAuthenticator{
		ApiKey: "HXoxv39NR1RyhzCxsJYhNCm9tFKbL1kni0mmnpJW7Fn5",
	}

	options := &assistantv2.AssistantV2Options{
		Version:       &ver,
		Authenticator: authenticator,
	}

	assistant, assistantErr := assistantv2.NewAssistantV2(options)

	if assistantErr != nil {
		panic(assistantErr)
	}

	assistant.SetServiceURL("https://api.eu-gb.assistant.watson.cloud.ibm.com")

	result, _, responseErr := assistant.CreateSession(assistant.
		NewCreateSessionOptions("1e108d47-e09f-4b70-b0c4-a8b1e0ac1813"))
	if responseErr != nil {
		panic(responseErr)
	}
	fmt.Println(*result.SessionID)
}
