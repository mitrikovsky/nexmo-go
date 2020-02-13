package nexmo

import "github.com/dghubble/sling"

// Work with the SMS API to send SMS messges. More information about this API: https://developer.nexmo.com/messaging/sms
type SMSService struct {
	sling   *sling.Sling
	authSet *AuthSet
}

type MessageType string

const (
	MessageTypeText    MessageType = "text"
	MessageTypeBinary  MessageType = "binary"
	MessageTypeWAPPush MessageType = "wappush"
	MessageTypeUnicode MessageType = "unicode"
)

func newSMSService(base *sling.Sling, authSet *AuthSet) *SMSService {
	sling := base.Base("https://rest.nexmo.com/sms/")
	return &SMSService{
		sling:   sling,
		authSet: authSet,
	}
}

func (c *SMSService) SetBaseURL(baseURL string) {
	c.sling.Base(baseURL)
}
