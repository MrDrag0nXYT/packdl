package util

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var ErrEmptyConfigHashsum = errors.New("File hashsum in config is empty")

func GetFileHash(filePath string) (string, error) {
	hasher := sha1.New()

	targetFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return "", err
	}
	defer targetFile.Close()

	if _, err = io.Copy(hasher, targetFile); err != nil {
		fmt.Println("Error while calculating hashsum of file:", err)
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func VerifyHash(filePath string, sha1Hash string) (bool, error) {
	if sha1Hash == "" {
		return false, ErrEmptyConfigHashsum
	}

	fileHash, err := GetFileHash(filePath)
	if err != nil {
		return false, err
	}

	// fmt.Println(" | Config hash:", sha1Hash)
	// fmt.Println(" | File hash:  ", fileHash)

	return strings.EqualFold(fileHash, sha1Hash), nil
}
