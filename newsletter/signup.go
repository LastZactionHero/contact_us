package newsletter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type payloadMCSignup struct {
	Email       string `json:"email_address"`
	Status      string `json:"status"`
	StatusIfNew string `json:"status_if_new"`
}

// Signup for Mailchimp newsletter
func Signup(email string) {
	mcAPIHost := os.Getenv("MAILCHIMP_DC_HOST")
	mcAPIKey := os.Getenv("MAILCHIMP_API_KEY")
	mcListID := os.Getenv("MAILCHIMP_CONTRACTOR_SIGNUP_LIST_ID")

	payload := payloadMCSignup{
		Email:       email,
		Status:      "subscribed",
		StatusIfNew: "subscribed",
	}
	jsonBytes, _ := json.Marshal(payload)

	requestPath := fmt.Sprintf("%s/lists/%s/members", mcAPIHost, mcListID)

	request, err := http.NewRequest("POST", requestPath, bytes.NewBuffer(jsonBytes))
	request.SetBasicAuth("username", mcAPIKey)

	fmt.Println("Sending")
	client := &http.Client{}
	response, err := client.Do(request)
	fmt.Println("Sent")
	if err != nil {
		fmt.Print(err)
	}
	if response.StatusCode != http.StatusOK {
		fmt.Print(response)
	}
}
