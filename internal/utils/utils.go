package utils

import (
	"encoding/base64"
	"regexp"
	"strings"
	"time"

	"github.com/dongri/phonenumber"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// RandomID to generate random uuid.
func RandomID() string {
	id := uuid.New()
	return strings.Replace(id.String(), "-", "", -1)
}

// CleanPhone to clean phone number.
func CleanPhone(str string, country string) string {
	return phonenumber.ParseWithLandLine(str, country)
}

// BoolToPtr to create boolean pointer from boolean.
func BoolToPtr(b bool) *bool {
	return &b
}

// TimeToPtr to create time pointer from time.
func TimeToPtr(t time.Time) *time.Time {
	return &t
}

// CleanSpace to remove space in string.
func CleanSpace(str string) string {
	return strings.Replace(str, " ", "", -1)
}

// RoundUp to round up float.
func RoundUp(v float64) int {
	if v != float64(int(v)) {
		return int(v) + 1
	}
	return int(v)
}

// AlphaOnly to convert to alpha string.
func AlphaOnly(str string) string {
	str = regexp.MustCompile(`[^a-zA-Z\s]+`).ReplaceAllString(str, "")
	str = regexp.MustCompile(`\s+`).ReplaceAllString(str, " ")
	return strings.TrimSpace(str)
}

// Truncate to truncate string.
func Truncate(str string, l int) string {
	if len(str) <= l {
		return str
	}
	return str[:l]
}

// ParseBasicAuth to parse basic auth.
func ParseBasicAuth(auth string) (username string, password string, ok bool) {
	authSplit := strings.Split(auth, " ")
	if len(authSplit) != 2 {
		return
	}
	if strings.ToLower(authSplit[0]) != "basic" {
		return
	}
	c, err := base64.StdEncoding.DecodeString(authSplit[1])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}

// TimePtrToProto to convert pointer time to timestamp proto.
func TimePtrToProto(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

// Repeat to repeat string.
func Repeat(str string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(str, n)
}

// PadRight to pad string to the right.
func PadRight(str string, l int, p string) string {
	return str + Repeat(p, l-len([]rune(str)))
}

// InArray to check if value is in array.
func InArray(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
