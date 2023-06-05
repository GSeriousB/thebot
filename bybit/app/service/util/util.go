package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(password string, hashedPassword string) bool {
	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func GenerateHash(value string) (string, error) {
	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ValidateEmail(Email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(Email)
}

// Bool stores v in a new bool value and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int32 stores v in a new int value and returns a pointer to it.
func Int(v int) *int { return &v }

// Int32 stores v in a new int32 value and returns a pointer to it.
func Int32(v int32) *int32 { return &v }

// Int64 stores v in a new int64 value and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// Float32 stores v in a new float32 value and returns a pointer to it.
func Float32(v float32) *float32 { return &v }

// Float64 stores v in a new float64 value and returns a pointer to it.
func Float64(v float64) *float64 { return &v }

// Uint32 stores v in a new uint32 value and returns a pointer to it.
func Uint32(v uint32) *uint32 { return &v }

// Uint64 stores v in a new uint64 value and returns a pointer to it.
func Uint64(v uint64) *uint64 { return &v }

// String stores v in a new string value and returns a pointer to it.
func String(v string) *string { return &v }

// Array stores v in a new slice and returns a pointer to it.
func Array(v []string) *[]string { return &v }

func UUID(v uuid.UUID) *uuid.UUID { return &v }

func IsNull(v interface{}, t string) bool {
	if t == "string" && v.(string) == "" {
		return true
	} else if t == "int" && v == 0 {
		return true
	} else if t == "datetime" && v.(time.Time).IsZero() {
		return true
	}
	return false
}

func IsValidDataType(dataType string, v interface{}) bool {
	// actualDataType := GetDataType(v)
	// if dataType == "number" {
	// 	if actualDataType == "int32" {
	// 		return true
	// 	} else if actualDataType == "int64" {
	// 		return true
	// 	} else if actualDataType == "float32" {
	// 		return true
	// 	} else if actualDataType == "float64" {
	// 		return true
	// 	} else if actualDataType == "uint32" {
	// 		return true
	// 	} else if actualDataType == "uint64" {
	// 		return true
	// 	}
	// 	return false
	// }
	return GetDataType(v) == dataType
}

func GetDataType(v interface{}) string {
	switch v.(type) {
	// Boolean
	case bool:
		return "bool"
	// Number
	case int32:
		return "int32"
	case int64:
		return "int64"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case uint32:
		return "uint32"
	case uint64:
		return "uint64"
	// String
	case string:
		return "string"
	}
	return ""
}

func IsContain(str []string, sub string) bool {
	for _, v := range str {
		if v == sub {
			return true
		}
	}
	return false
}

func InterfaceStringToArray(claims interface{}) []string {
	aInterface := claims.([]interface{})
	aString := make([]string, len(aInterface))
	for i, v := range aInterface {
		aString[i] = v.(string)
	}
	return aString
}

// TODO: This can break in case the role_code has duplicates. Ex - ["appuser", "company-admin", "appuser"]
func Subset(availArray, subArray []string) bool {
	set := make(map[string]bool)
	for _, value := range availArray {
		set[value] = true
	}
	f := 0
	for _, value := range subArray {
		if _, found := set[value]; found {
			f = f + 1
		}
	}
	return f == len(subArray)
}

func RemoveEmptyStringFromArray(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func GenerateSignature(msg, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)

	return hex.EncodeToString(mac.Sum(nil))
}

func VerifySignature(msg, key []byte, hash string) (bool, error) {
	sig, err := hex.DecodeString(hash)
	if err != nil {
		return false, err
	}

	mac := hmac.New(sha256.New, key)
	mac.Write(msg)

	return hmac.Equal(sig, mac.Sum(nil)), nil
}
