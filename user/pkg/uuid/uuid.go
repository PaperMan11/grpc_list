package uuid

import (
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letterWithNumberAndSpaceBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

/*
	produce random strings with the English character only,
	length ranging from min to max
*/
func RandStringBytes(min, max int) []byte {
	if min >= max || max == 0 || min == 0 {
		return nil
	}

	n := rand.Intn(max-min) + min
	b := make([]byte, n)
	for i := 0; i < n-1; i++ {
		b[i] = letterWithNumberAndSpaceBytes[rand.Intn(len(letterWithNumberAndSpaceBytes))]
	}
	b[n-1] = 0
	return b
}

const UIDLen = 36

// UUIDWithLen 指定长度UUID
func UUIDWithLen(length int) string {
	//uuid := uuid.NewV4().String()
	uuid := uuid.NewV4().String()
	// uuid = strings.Replace(uuid, "-", "", -1)
	return uuid[:length]
}

// UUID
func UUID() string {
	return UUIDWithLen(UIDLen)
}
