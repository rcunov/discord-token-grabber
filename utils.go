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

// This function checks if an inputted string is empty.
// If so, then it returns a Discord "no" emote. If the
// string is not empty, then it returns the input string.
func nullStringEmote(input string) (output string) {
	if input == "" {
		return ":prohibited:"
	} else {
		return input
	}
}

// This function checks if an inputted bool is true.
// If so, then it returns a Discord "check mark" emote.
// If the bool is not true, then it returns a "no" emote.
func boolToEmote(input bool) (output string) {
	if input {
		return ":white_check_mark:"
	} else {
		return ":prohibited:"
	}
}