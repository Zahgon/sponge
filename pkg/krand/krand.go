// Package krand is a library for generating random strings, integers, floating point numbers.
package krand

import (
	"math/rand"
	"time"
)

// nolint
const (
	R_NUM   = 1 // only number
	R_UPPER = 2 // only capital letters
	R_LOWER = 4 // only lowercase letters
	R_All   = 7 // numbers, upper and lower case letters
)

var (
	refSlices = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	kinds     = [][]byte{refSlices[0:10], refSlices[10:36], refSlices[0:36], refSlices[36:62], refSlices[36:], refSlices[10:62], refSlices[0:62]}
)

func init() {
	rand.Seed(time.Now().UnixNano()) //nolint
}

// String generate random strings of any length of multiple types, default length is 6 if size is empty
// example: String(R_ALL), String(R_ALL, 16), String(R_NUM|R_LOWER, 16)
func String(kind int, size ...int) string { _ = "STUB: not implemented"; return "" }

// Bytes generate random strings of any length of multiple types, default length is 6 if bytesLen is empty
// example: Bytes(R_ALL), Bytes(R_ALL, 16), Bytes(R_NUM|R_LOWER, 16)
func Bytes(kind int, bytesLen ...int) []byte { _ = "STUB: not implemented"; return nil }

// default length 6

// Int generate random numbers of specified range size,
// compatible with Int(), Int(max), Int(min, max), Int(max, min) 4 ways, min<=random number<=max
func Int(rangeSize ...int) int { _ = "STUB: not implemented"; return 0 }

// default 0~100

// Float64 generates a random floating point number of the specified range size,
// Four types of passing references are supported, example: Float64(dpLength), Float64(dpLength, max),
// Float64(dpLength, min, max), Float64(dpLength, max, min), min<=random numbers<=max
func Float64(dpLength int, rangeSize ...int) float64 { _ = "STUB: not implemented"; return 0 }

// default 0~100

// NewID Generate a milliseconds+random number ID.
func NewID() int64 { _ = "STUB: not implemented"; return 0 }

// NewStringID Generate a string ID, the hexadecimal form of NewID(), total 16 bytes.
func NewStringID() string { _ = "STUB: not implemented"; return "" }

// NewSeriesID Generate a datetime+random string ID,
// datetime is microsecond precision, 20  bytes, random is 6 bytes, total 26 bytes.
// example: 20060102150405000000123456
func NewSeriesID() string {
	_ = "STUB: not implemented"
	// Declare a buffer, only 26 bytes are needed
	return ""
}

// Format datetime with microsecond precision, and store in the buffer

// Generate a 6-digit random number and append it to the buffer

// Return the final string without the dot
