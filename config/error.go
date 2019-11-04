package config

import "errors"

var (
	ErrDataNotFound = errors.New("the requested resource doesn't exists")
)
