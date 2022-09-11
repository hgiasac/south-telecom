package sms

import (
	"net/http"
	"net/url"
)

// SendMessageInput represents the message request input
type SendMessageInput struct {
	// Brand name of the message
	// Represents sender ID and it can be alphanumeric or numeric.
	// Alphanumeric sender ID length should be between 3 and 11 characters
	From string `json:"from"`
	// Message destination address.
	// Destination address must be in international format (example: 84987654321)
	To string `json:"to"`
	// Text of the message that will be sent.
	Text string `json:"text"`
	// send message content with or without unicode
	Unicode TextType `json:"unicode"`
	// If dlr=1, report form telco will sent to customer via API
	ReportType ReportType `json:"dlr"`
	// ID of the message, defined by the customer.
	SmsID string `json:"smsid,omitempty"`
	// Campaign ID for management
	CampaignID string `json:"messageid,omitempty"`
	// Whether sending encrypted message
	Encrypted EncryptionType `json:"encrypted"`
	// Content type of the message
	ContentID ContentType `json:"contentid"`
	// Must require encrypting the message to Viettel Bank or not
	MustEncryptViettel ViettelEncryptionType `json:"mustencryptviettel"`
}

// SendMessageResponse represents the send message response
type SendMessageResponse struct {
	Status    ResponseStatus `json:"status"`
	ErrorCode ErrorCode      `json:"errorcode,omitempty"`
	// Mobile number portability status
	// The subscriber may register to switch network but hold numbers
	MNP MobileNumberPortability `json:"mnp"`
	// the current telecom provider of the subscriber
	Carrier string `json:"carrier"`
}

type smsService struct {
	client *httpClient
}

// SendMessage send SMS message to the target phone number
func (ss *smsService) SendMessage(input SendMessageInput) (*SendMessageResponse, *http.Response, error) {
	u, err := url.Parse("/webapi/sendSMS")
	if err != nil {
		return nil, nil, err
	}
	// create the request
	req, err := ss.client.NewRequest("POST", u.String(), input)
	if err != nil {
		return nil, nil, err
	}

	result := &SendMessageResponse{}
	resp, err := ss.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, err
}
