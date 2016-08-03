package getstream

// import (
// 	"crypto/hmac"
// 	"crypto/sha1"
// 	"encoding/base64"
// 	"strings"
// )
//
// func urlSafe(src string) string {
// 	src = strings.Replace(src, "+", "-", -1)
// 	src = strings.Replace(src, "/", "_", -1)
// 	src = strings.Trim(src, "=")
// 	return src
// }
//
// func Sign(secret, message string) string {
// 	hash := sha1.New()
// 	hash.Write([]byte(secret))
// 	key := hash.Sum(nil)
// 	mac := hmac.New(sha1.New, key)
// 	mac.Write([]byte(message))
// 	digest := base64.StdEncoding.EncodeToString(mac.Sum(nil))
// 	return urlSafe(digest)
// }
//
// func SignSlug(secret string, slug Slug) Slug {
// 	return Slug{slug.Slug, slug.ID, Sign(secret, slug.Slug+slug.ID)}
// }
//
// func SignActivity(secret string, activity *Activity) *Activity {
// 	result := &Activity{}
// 	*result = *activity
// 	for i, slug := range result.To {
// 		result.To[i] = SignSlug(secret, slug)
// 	}
//
// 	return result
// }
