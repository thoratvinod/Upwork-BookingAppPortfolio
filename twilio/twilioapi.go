package twilioapi

import (
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSRequest struct {
	From string
	To   string
	Body string
}

func SendSMS(sms *SMSRequest) error {
	// TWILIO_ACCOUNT_SID and TWILIO_AUTH_TOKEN should be exists in env
	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetBody(sms.Body)
	params.SetFrom(sms.From)
	params.SetTo(sms.To)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		return err
	}
	return nil
}
