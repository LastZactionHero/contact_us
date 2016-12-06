package newsletter

import (
	"encoding/json"
	"os"
)

type payloaMCSignup struct {
	Email       string `json:"email_address"`
	Status      string `json:"status"`
	StatusIfNew string `json:"status_if_new"`
}

func signup(email) {
	mcAPIHost := os.Getenv("MAILCHIMP_DC_HOST")
	mcAPIKey := os.Getenv("MAILCHIMP_API_KEY")
	mcListID := os.Getenv("MAILCHIMP_CONTRACTOR_SIGNUP_LIST_ID")

	requestPath := fmt.Spritnf("%s/lists/%s/members", mcAPIHost, mcListID)
	json.Marshal(payloaMCSignup)
}

// MAILCHIMP_API_KEY
// MAILCHIMP_CONTRACTOR_SIGNUP_LIST_ID
// MAILCHIMP_DC_HOST

// newsletter id? 76705
//
// POST lists/83513a6597/members
// https://us14.api.mailchimp.com/3.0/lists/{list_id}.
//
//
//
// {
// 	"email_address": "zach+mc1@squarewaveng.com",
// 	"status": "subscribed",
// 	"status_if_new": "subscribed"
// }
