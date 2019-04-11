package docker

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
)

func Checksum(filename string) (string, error) {
	hasher := sha256.New()
	s, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	hasher.Write(s)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
