package gofirebaseauthingcpfunctions

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	firebaseAuth "firebase.google.com/go/auth"
)

// AuthenticateFirebaseUser extracts firebase user email from the http request.
// Returns *firebase.google.com/go/auth.UserRecord if authenticated, error if not authenticated
func AuthenticateFirebaseUser(ctx context.Context, r *http.Request, authClient *firebaseAuth.Client) (*firebaseAuth.UserRecord, error) {
	tokenStr, err := getBearerToken(r)
	if err != nil {
		return nil, err
	}
	authToken, err := authClient.VerifyIDTokenAndCheckRevoked(ctx, tokenStr)
	if err != nil {
		return nil, fmt.Errorf("Canot verify ID token: %v", err)
	}

	uid := authToken.UID
	if uid == "" {
		return nil, fmt.Errorf("ID token verification failed")
	}

	userRecord, err := authClient.GetUser(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("Canot get firebase user: %v", err)
	}

	return userRecord, nil
}

func getBearerToken(r *http.Request) (string, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return "", fmt.Errorf("Empty authorization header")
	}

	headerParts := strings.SplitN(authorizationHeader, " ", 2)
	if len(headerParts) < 2 {
		return "", fmt.Errorf("Invalid authorization header")
	}

	tokenType := headerParts[0]
	token := headerParts[1]

	if strings.ToLower(tokenType) != "bearer" {
		return "", fmt.Errorf("Invalid authorization token type (%s). Only 'bearer' wllowed", tokenType)
	}

	return token, nil
}
