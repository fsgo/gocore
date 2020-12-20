// +build go1.13

/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/12/20
 */

package gocore

import (
	"encoding/json"
	"io"
	"mime/multipart"
)

// MultiPartReader interface off *multipart.Reader
type MultiPartReader interface {
	NextPart() (*multipart.Part, error)
	ReadForm(maxMemory int64) (*multipart.Form, error)
}

// JSONDecoder interface of *json.Decoder
type JSONDecoder interface {
	Buffered() io.Reader
	Decode(v interface{}) error
	DisallowUnknownFields()
	More() bool
	Token() (json.Token, error)
	UseNumber()
}
