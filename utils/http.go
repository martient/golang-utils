package utils

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

func NewHTTPOAuth2Client(ctx context.Context, accessToken string) *http.Client {
	if accessToken == "" {
		return http.DefaultClient
	}
	src := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	return oauth2.NewClient(ctx, src)
}
