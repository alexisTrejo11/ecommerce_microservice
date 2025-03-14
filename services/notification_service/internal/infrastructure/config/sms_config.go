package config

import (
	"log"
	"os"
)

type SMSConfig struct {
	AccountSid   string
	AuthToken    string
	TwilioNumber string
}

func NewSMSConfig() *SMSConfig {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	twilioNumber := os.Getenv("TWILIO_PHONE_NUMBER")

	if accountSid == "" || authToken == "" || twilioNumber == "" {
		log.Fatalf("Can't provide SMS service: missing Twilio configuration")
	}

	return &SMSConfig{
		AccountSid:   accountSid,
		AuthToken:    authToken,
		TwilioNumber: twilioNumber,
	}
}
