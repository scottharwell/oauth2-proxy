package providers

import (
	"context"
	"fmt"

	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/apis/sessions"
	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/requests"
)

// OIDCProvider represents an OIDC based Identity Provider
type IDCSProvider struct {
	*ProviderData
}

// NewOIDCProvider initiates a new OIDCProvider
func NewIDCSProvider(p *ProviderData) *IDCSProvider {
	p.ProviderName = "Oracle Identity Cloud"
	return &IDCSProvider{ProviderData: p}
}

var _ Provider = (*IDCSProvider)(nil)

// GetEmailAddress returns the Account email address
func (p *IDCSProvider) GetEmailAddress(ctx context.Context, s *sessions.SessionState) (string, error) {
	json, err := requests.New(p.ValidateURL.String()).
		WithContext(ctx).
		WithHeaders(makeOIDCHeader(s.AccessToken)).
		Do().
		UnmarshalJSON()
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}

	email, err := json.Get("sub").String()
	return email, err
}
