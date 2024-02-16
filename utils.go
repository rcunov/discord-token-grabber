package main

import (
	"encoding/base64"
)

// This function Base64 encodes a string.
// Input is a string to be encoded.
// Output is the Base64 encoded string.
func base64EncodeStr(str string) string {
    return base64.StdEncoding.EncodeToString([]byte(str))
}