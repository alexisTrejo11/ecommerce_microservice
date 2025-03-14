package sms

import (
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/infrastructure/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSService struct {
	Client *twilio.RestClient
	From   string
}

func NewSMSService(sms *config.SMSConfig) *SMSService {
	client := twilio.NewRestClient()
	client.SetRegion("au1")
	client.SetEdge("sydney")

	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: sms.AccountSid,
		Password: sms.AuthToken,
		Client:   client.Client,
	})

	return &SMSService{
		Client: twilioClient,
		From:   sms.TwilioNumber,
	}
}

func (sms *SMSService) SendSMS(phoneNumber, message string, twilioNumber string) (string, error) {
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(twilioNumber)
	params.SetBody(message)

	resp, err := sms.Client.Api.CreateMessage(params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}
