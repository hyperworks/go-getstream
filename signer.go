package getstream

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

// Signer is responsible for generating Tokens
type Signer struct {
	Secret string
}

func (s Signer) urlSafe(src string) string {
	src = strings.Replace(src, "+", "-", -1)
	src = strings.Replace(src, "/", "_", -1)
	src = strings.Trim(src, "=")
	return src
}

// generateToken will user the Secret of the signer and the message passed as an argument to generate a Token
func (s Signer) generateToken(message string) string {
	hash := sha1.New()
	hash.Write([]byte(s.Secret))
	key := hash.Sum(nil)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(message))
	digest := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return s.urlSafe(digest)
}
