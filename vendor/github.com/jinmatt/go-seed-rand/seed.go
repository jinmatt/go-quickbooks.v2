package seed

import (
	randc "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"math/rand"
	"sync"
	"time"
)

// RandomKeyCharacters is a []byte of the characters to choose from when generating
// random keys.
var RandomKeyCharacters = []byte("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var RandomKeyIntegers = []byte("0123456789")

var randomOnce sync.Once

func init() {
	seed := time.Now().UnixNano()
	b, err := GenerateRandomBytes(64)
	if err == nil {
		seed = int64(binary.BigEndian.Uint64(b))
	}
	rand.Seed(seed)
}

// RandomKey generates a random key at the given length.
//
// The first time this is called, the rand.Seed will be set
// to the current time.
func RandomKey(length int) string {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		randInt := randInt(0, len(RandomKeyCharacters))
		bytes[i] = RandomKeyCharacters[randInt : randInt+1][0]
	}
	return string(bytes)
}

// randInt generates a random integer between min and max.
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// RandomIntKey generates a random key at the given length.
//
// The first time this is called, the rand.Seed will be set
// to the current time.
func RandomIntKey(length int) string {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		randInt := randInt(0, len(RandomKeyIntegers))
		bytes[i] = RandomKeyIntegers[randInt : randInt+1][0]
	}
	return string(bytes)
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := randc.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
