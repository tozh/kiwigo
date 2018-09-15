package client

import (
	"errors"
	"io"
)

var errTimeOut = errors.New("db server time out")
var errSetFailed = errors.New("set failed")
var errEOF = io.EOF
var errWrongType = errors.New("the value of the key has the wrong type")
var errNil = errors.New("the value does not exist")
var errWrongFormat = errors.New("wrong format")