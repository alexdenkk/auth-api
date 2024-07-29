package hash

import (
	"crypto/sha1"
	"encoding/base64"
)

// Hash - SHA1 hashing function
func Hash(text string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
