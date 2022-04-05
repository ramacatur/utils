package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"regexp"
)

//MD5Hash for encode string to md5
func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

//Sha1 for encode string to sha1
func Sha1(input string) string {
	if input == "" {
		return "adc83b19e793491b1c6ea0fd8b46cd9f32e592fc"
	}
	return fmt.Sprintf("%x", sha1.Sum([]byte(input)))
}

//Sha256 for encode string to sha256
func Sha256(input string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
}

//HMAC256 for encode string to hmac-sha256
func HMAC256(secret, message string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))

	//base64.StdEncoding.EncodeToString(h.Sum(nil))
	return sha
}

//Secret2Password for generate password
func Secret2Password(username, secret string) string {
	return Sha1(Sha1(secret[:8]) + Sha1(username) + Sha1(secret[8:]))
}

//Base64Encode func
func Base64Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

//Base64Decode fuunc
func Base64Decode(input string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(input)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

// IsValidUUID for check the string is valid UUID V4 or not
func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
