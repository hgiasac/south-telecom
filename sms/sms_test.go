package sms

import (
	"fmt"
	"log"
	"testing"
)

func setup(t *testing.T) *Client {
	apiKey := "random_secret"

	client, err := NewClient(apiKey)
	if err != nil {
		t.Fatal(err)
	}
	client.logger = func(s string, i ...interface{}) {
		log.Println(s, i)
	}
	return client
}

func TestSendMessage(t *testing.T) {
	client := setup(t)

	brandName := "VienThongMN"
	result, _, err := client.Sms.SendMessage(SendMessageInput{
		From:               brandName,
		To:                 "84932123456",
		Text:               fmt.Sprintf("%s gui ma OTP dang nhap ung dung Dich Vu %s cua ban la 295788. Ma xac minh co hieu luc trong 3 phut.", brandName, brandName),
		Unicode:            NewTextAscii(),
		ReportType:         NoReport,
		SmsID:              "random_id",
		CampaignID:         "random_campaign_id",
		Encrypted:          NonEncrypted,
		ContentID:          NormalMessage,
		MustEncryptViettel: EncryptedOnDemand,
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Status != StatusFailure {
		t.Fatalf("Status: expected %d, got %d", StatusFailure, result.Status)
	}
	if result.ErrorCode != ErrCodeUnauthorized {
		t.Fatalf("ErrorCode: expected %d, got %d", ErrCodeUnauthorized, result.ErrorCode)
	}
}

func TestSendMessageDebit(t *testing.T) {
	client := setup(t)

	brandName := "VienThongMN"
	result, _, err := client.Sms.SendMessageDebit(SendMessageDebitInput{
		From:       brandName,
		To:         "84932123456",
		Text:       fmt.Sprintf("%s gui ma OTP dang nhap ung dung Dich Vu %s cua ban la 295788. Ma xac minh co hieu luc trong 3 phut.", brandName, brandName),
		ReportType: NoReport,
		SmsID:      "random_id",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Status != StatusFailure {
		t.Fatalf("Status: expected %d, got %d", StatusFailure, result.Status)
	}
	if result.ErrorCode != ErrCodeUnauthorized {
		t.Fatalf("ErrorCode: expected %d, got %d", ErrCodeUnauthorized, result.ErrorCode)
	}
}
