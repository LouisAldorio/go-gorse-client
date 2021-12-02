package gorseclient

import (
	"errors"
)

var (
	Err404NotFound            = errors.New("404 not found")
	Err500InternalSystemError = errors.New("500 internal system error")
)