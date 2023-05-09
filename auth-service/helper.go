package main

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateField(field string, value string, required bool, pattern *regexp.Regexp) error {
	isEmpty := strings.TrimSpace(value) == ""
	if required && isEmpty {
		return errors.New(field + " cannot be empty")
	}

	if pattern != nil && !isEmpty && !pattern.MatchString(value) {
		return errors.New(field + " is not valid")
	}

	return nil
}
