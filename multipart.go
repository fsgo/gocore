/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/12/20
 */

package gocore

import (
	"io"
	"mime/multipart"
	"net/textproto"
)

// MultiPartReader interface off *multipart.Reader
type MultiPartReader interface {
	NextPart() (*multipart.Part, error)
	NextRawPart() (*multipart.Part, error)
	ReadForm(maxMemory int64) (*multipart.Form, error)
}

var _ MultiPartReader = (*multipart.Reader)(nil)

// MultiPartWriter interface off *multipart.Writer
type MultiPartWriter interface {
	Boundary() string
	Close() error
	CreateFormField(fieldname string) (io.Writer, error)
	CreateFormFile(fieldname, filename string) (io.Writer, error)
	CreatePart(header textproto.MIMEHeader) (io.Writer, error)
	FormDataContentType() string
	SetBoundary(boundary string) error
	WriteField(fieldname, value string) error
}

var _ MultiPartWriter = (*multipart.Writer)(nil)

// MultiPartPart interface off *multipart.Part
type MultiPartPart interface {
	Close() error
	FileName() string
	FormName() string
	Read(d []byte) (n int, err error)
}

var _ MultiPartPart = (*multipart.Part)(nil)
