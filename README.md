# go-firebase-auth-in-gcp-functions
Verify firebase auth in your golang GCP cloud functions

## Installation

```bash
$ go get github.com/Jblew/go-firebase-auth-in-gcp-functions@1.0.0
```

## Usage

```go
package functions

import (
  "context"
  "fmt"
  "log"
  "net/http"

  firebaseGcpAuth "github.com/Jblew/go-firebase-auth-in-gcp-functions"
  auth "firebase.google.com/go/auth"
)

func SomeGCPHttpCloudFunction(w http.ResponseWriter, req *http.Request) error {
   // You need to provide 1. Context, 2. request, 3. firebase auth client
  var client *auth.Client
  firebaseUser, err := firebaseGcpAuth.AuthenticateFirebaseUser(context.Background(), req, authClient)
  if err != nil {
    return err // Error if not authenticated or bearer token invalid
  }

  // Returned value:
  var firebaseUser *auth.UserRecord
}

```
