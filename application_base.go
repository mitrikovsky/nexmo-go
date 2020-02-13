package nexmo

import "github.com/dghubble/sling"

// Use the Application API to create and manage your applications. More info: https://developer.nexmo.com/application
type ApplicationService struct {
	sling   *sling.Sling
	authSet *AuthSet
}

func newApplicationService(base *sling.Sling, authSet *AuthSet) *ApplicationService {
	sling := base.Base("https://api.nexmo.com/v1/applications/")
	return &ApplicationService{
		sling:   sling,
		authSet: authSet,
	}
}

// Set the base URL for the API calls. Useful for testing.
func (c *ApplicationService) SetBaseURL(baseURL string) {
	c.sling.Base(baseURL)
}
