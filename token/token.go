package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

const (
	fieldDelimiter = "~"
)

// EdgeAuthToken can be used to generate an authentication token for Akamai URLs
type EdgeAuthToken struct {
	Key           string
	WindowSeconds int
	ClientIP      string
}

// GenerateURLToken generates a token to be used to authenticate to Akamai
// url = The path portion of the url being requested (eg, /my/asset)
func (e EdgeAuthToken) GenerateURLToken(url string) (string, error) {
	windowSeconds := e.WindowSeconds
	if e.WindowSeconds == 0 {
		windowSeconds = 300
	}
	startTime := time.Now().Unix()
	endTime := startTime + int64(windowSeconds)

	var newToken []string
	if e.ClientIP != "" {
		newToken = append(newToken, fmt.Sprintf("ip=%s", e.ClientIP))
	}

	//new_token = append(new_token, fmt.Sprintf("st=%d",start_time))
	newToken = append(newToken, fmt.Sprintf("exp=%d", endTime))

	hashSource := make([]string, len(newToken))
	copy(hashSource, newToken)

	hashSource = append(hashSource, fmt.Sprintf("url=%s", url))

	key, err := hex.DecodeString(e.Key)
	if err != nil {
		return "", err
	}

	mac := hmac.New(sha256.New, key)
	data := []byte(strings.Join(hashSource, fieldDelimiter))

	mac.Write(data)
	sum := strings.ToLower(hex.EncodeToString(mac.Sum(nil)))
	newToken = append(newToken, fmt.Sprintf("hmac=%s", sum))

	return strings.Join(newToken, fieldDelimiter), nil
}
