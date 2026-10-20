package utils

import "errors"

var ErrEmailExists = errors.New("email already registered")
var ErrEmailNotExists = errors.New("email not exist")
var ErrUrlExists = errors.New("Url already exists")
