package encrypt

import (
	"encoding/base64"
	"fmt"
)

// Base64Encoding
func Base64Encoding(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decoding
func Base64Decoding(str string) (dec []byte){
	dec, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Printf("Error decoding string: %s", err.Error())
		return
	}
	return
}

// Base64UrlEncoding
// URL and filename safe encoding and encoding
func Base64UrlEncoding(str string) string {
	return base64.URLEncoding.EncodeToString([]byte(str))
}

// Base64UrlDecoding
// URL and filename safe encoding and decoding
func Base64UrlDecoding(str string) (dec []byte) {
	dec, err := base64.URLEncoding.DecodeString(str)
	if err != nil {
		fmt.Printf("Error decoding url string: %s", err.Error())
		return
	}
	return
}

// Base64RawEncoding
// No padding encoding
func Base64RawEncoding(str string) string {
	return base64.RawStdEncoding.EncodeToString([]byte(str))
}

// Base64RawDecoding
// // No padding decoding
func Base64RawDecoding(str string) (dec []byte) {
	dec, err := base64.RawStdEncoding.DecodeString(str)
	if err != nil {
		fmt.Printf("Error decoding no padding err: %s", err.Error())
	}
	return dec
}

// Base64RawUrlEncoding
// Base64 Url Encoding without Padding
func Base64RawUrlEncoding(str string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(str))
}

// Base64RowUrlDecoding
// Base64 Url Decoding without Padding
func Base64RowUrlDecoding(str string) (dec []byte) {
	dec, err := base64.RawURLEncoding.DecodeString(str)
	if err != nil {
		fmt.Printf("Error url decoding without padding")
		return
	}
	return
}
