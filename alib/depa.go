package alib

import (
	"errors"
)

func GetPageContent() (string, error) {
	err := errors.New("simple GetPageContent error")
	pageContent := "<head></head><body></body>"

	return pageContent, err
}
