package utils

import (
	"context"
	"errors"

	"main/internal/config"
	"google.golang.org/api/idtoken"
)

// VerifyGoogleIDToken verifies a Google ID token using the official Google SDK.
// Returns email, google user id (sub), and name.
func VerifyGoogleIDToken(idToken string) (email, sub, name string, err error) {
	clientID := config.GetEnv("GOOGLE_CLIENT_ID", "")
	if clientID == "" {
		return "", "", "", errors.New("missing GOOGLE_CLIENT_ID")
	}

	payload, err := idtoken.Validate(context.Background(), idToken, clientID)
	if err != nil {
		return "", "", "", err
	}

	// Extract claims
	if v, ok := payload.Claims["email"].(string); ok {
		email = v
	}
	if v, ok := payload.Claims["sub"].(string); ok {
		sub = v
	}
	if v, ok := payload.Claims["name"].(string); ok {
		name = v
	}
	return
}
