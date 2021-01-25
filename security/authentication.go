package security

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/url"
	"time"
)

type AuthenticationHeaders struct {
	Token     string `json:"token"`
	Date      string `json:"date"`
	Signature string `json:"signature"`
}

func GenerateAuthenticationHeaders(publicToken, privateToken, requestPath string, params *url.Values) *AuthenticationHeaders {
	currentTime := time.Now().UTC().Format(time.RFC3339)

	return &AuthenticationHeaders{
		Token:     publicToken,
		Date:      currentTime,
		Signature: calculateSignature(publicToken, privateToken, requestPath, currentTime, params),
	}
}

func calculateSignature(publicToken, privateToken, requestPath, date string, params *url.Values) string {
	mac := hmac.New(sha1.New, []byte(privateToken))
	s := ""
	if params != nil {
		s = fmt.Sprintf("%s?%s\n%s\n%s", requestPath, params.Encode(), publicToken, date)
	} else {
		s = fmt.Sprintf("%s\n%s\n%s", requestPath, publicToken, date)
	}
	mac.Write([]byte(s))
	return hex.EncodeToString(mac.Sum(nil))
}
