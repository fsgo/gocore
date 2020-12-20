// +build !go1.13

/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/12/20
 */

package gocore

import (
	"mime/multipart"
)

// MultiPartReader interface off *multipart.Reader
type MultiPartReader interface {
	NextPart() (*multipart.Part, error)
	NextRawPart() (*multipart.Part, error)
	ReadForm(maxMemory int64) (*multipart.Form, error)
}
