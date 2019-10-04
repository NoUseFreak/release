package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func BuildURLHash(alg, url string) (string, error) {
	logrus.Tracef("Downloading %s", url)
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("Not found")
	}

	return BuildHash(alg, resp.Body)
}

func BuildHash(alg string, reader io.Reader) (string, error) {

	var hasher hash.Hash
	switch alg {
	case "sha256":
		hasher = sha256.New()
	case "sha1":
		hasher = sha1.New()
	default:
		return "", errors.New("Unknown hashing algorithm")
	}

	_, err := io.Copy(hasher, reader)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
