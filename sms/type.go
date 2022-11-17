package sms

type ResponseStatus int
type ErrorCode int

// each sms type has different phone number and cost
type FeeType int
type TextType int
type ReportType int
type EncryptionType int
type ViettelEncryptionType int
type ContentType int
type MobileNumberPortability int

const (
	StatusSuccess ResponseStatus = 1
	StatusFailure ResponseStatus = 0

	ErrCodeUnauthorized                    ErrorCode = 40
	ErrCodeInvalidPassword                 ErrorCode = 41
	ErrCodeInvalidUser                     ErrorCode = 42
	ErrCodeGatewayError                    ErrorCode = 50
	ErrCodeInvalidIP                       ErrorCode = 51
	ErrCodeInvalidInputParams              ErrorCode = 52
	ErrCodeInvalidPhoneNumber              ErrorCode = 53
	ErrCodeInvalidSender                   ErrorCode = 54
	ErrCodeInvalidContent                  ErrorCode = 55
	ErrCodeSendDebitSmsApiPermissionDenied ErrorCode = 80
	ErrCodeSendDebitSmsPermissionDenied    ErrorCode = 81
	ErrCodeAccountOverQuota                ErrorCode = 82
	ErrCodeInvalidMessageLength            ErrorCode = 551
	// The phone number is not registered to Viettel but the message is encrypted
	ErrCodeContentMustNotBeEncrypted ErrorCode = 552
	// The phone number is registered to Viettel but the message isn't encrypted
	ErrCodeContentMustBeEncrypted ErrorCode = 553
	// The subscriber has switched network and the new network is not registered with ST
	ErrCodeInvalidPhoneNumberMNP ErrorCode = 531

	NoFee  FeeType = 0
	HasFee FeeType = 1

	TextAscii   TextType = 0
	TextUnicode TextType = 1

	NoReport         ReportType = 0
	ReportToCustomer ReportType = 1

	NonEncrypted EncryptionType = 0
	Encrypted    EncryptionType = 1

	NormalMessage       ContentType = 0
	OTPMessage          ContentType = 1
	BalanceAlertMessage ContentType = 2

	EncryptedOnDemand     ViettelEncryptionType = 0
	MustEncryptForViettel ViettelEncryptionType = 1

	// Subscribers who are not registered to change network hold numbers
	SameNetwork MobileNumberPortability = 0
	// Subscribers who register to switch network hold numbers
	SwitchedNetwork MobileNumberPortability = 1
)

// NewTextAscii creates text ascii type reference
func NewTextAscii() *TextType {
	ty := TextAscii
	return &ty
}

// NewTextUnicode creates text unicode type reference
func NewTextUnicode() *TextType {
	ty := TextUnicode
	return &ty
}
