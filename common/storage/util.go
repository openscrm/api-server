package storage

import (
	"github.com/pkg/errors"
	"mime"
	"path/filepath"
	"regexp"
	"strings"
)

func GetContentType(filePath string) (contentType string, err error) {
	ext := strings.ToLower(filepath.Ext(filePath))
	if ext == "" {
		err = errors.New("file ext is required")
		return
	}

	contentType = mime.TypeByExtension(ext)
	if contentType == "" {
		err = errors.New("invalid file ext")
		return
	}

	return
}

func IsValidObjectKey(objectKey string) bool {
	return regexp.MustCompile("^[a-zA-Z0-9/\\-_]+\\.[a-zA-Z0-9]+$").MatchString(objectKey)
}
