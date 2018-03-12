package main

import (
	"fmt"
	"strings"

	"github.com/goware/urlx"
)

func validateURL(value string) (string, error) {
	value = strings.Trim(value, " ")
	if value == "" {
		return "", fmt.Errorf("Empty URL string")
	}
	u, err := urlx.Parse(value)
	if err != nil {
		return "", err
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return "", fmt.Errorf("Invalid scheme, please start URL with http:// or https://")
	}
	return u.String(), nil
}
