package service

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

func verifyHash(filePath string, sha1Hash string) (bool, error) {
	if sha1Hash == "" {
		return false, ErrEmptyConfigHashsum
	}

	hasher := sha1.New()

	targetFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return false, err
	}
	defer targetFile.Close()

	if _, err = io.Copy(hasher, targetFile); err != nil {
		fmt.Println("Error while calculating hashsum of file:", err)
		return false, err
	}

	fileHash := hex.EncodeToString(hasher.Sum(nil))

	// fmt.Println(" | Config hash:", sha1Hash)
	// fmt.Println(" | File hash:  ", fileHash)

	return strings.EqualFold(fileHash, sha1Hash), nil
}
