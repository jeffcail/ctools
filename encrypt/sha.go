package encrypt

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

// Sha1
func Sha1(str string) string {
	enc := sha1.New()
	io.WriteString(enc, str)
	return fmt.Sprintf("%v", enc.Sum(nil))
}

// Sha256
func Sha256(msg string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(msg))
	return fmt.Sprintf(base64.URLEncoding.EncodeToString(h.Sum(nil)))
}
