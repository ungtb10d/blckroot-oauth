package device_test

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/blckroot/oauth/device"
)

// This demonstrates how to perform OAuth Device Authorization Flow for GitHub.com.
// After RequestCode successfully completes, the client app should prompt the user to copy
// the UserCode and to open VerificationURI in their web browser to enter the code.
func ExampleRequestCode() {
	clientID := os.Getenv("OAUTH_CLIENT_ID")
	scopes := []string{"repo", "read:org"}
	httpClient := http.DefaultClient

	code, err := device.RequestCode(httpClient, "https://github.com/login/device/code", clientID, scopes)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Copy code: %s\n", code.UserCode)
	fmt.Printf("then open: %s\n", code.VerificationURI)

	accessToken, err := device.Wait(context.TODO(), httpClient, "https://github.com/login/oauth/access_token", device.WaitOptions{
		ClientID:   clientID,
		DeviceCode: code,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Access token: %s\n", accessToken.Token)
}
