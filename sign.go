package getstream

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

func urlSafe(src string) string {
	src = strings.Replace(src, "+", "-", -1)
	src = strings.Replace(src, "/", "_", -1)
	src = strings.Trim(src, "=")
	return src
}

func Sign(secret, feedId string) string {
	hash := sha1.New()
	hash.Write([]byte(secret))
	key := hash.Sum(nil)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(feedId))
	digest := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return urlSafe(digest)
}
