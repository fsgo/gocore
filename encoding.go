/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/12/20
 */

package gocore

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
)

// MarshalFunc encode
type MarshalFunc func(v interface{}) ([]byte, error)

// UnmarshalFunc decode
type UnmarshalFunc func(data []byte, v interface{}) error

// MarshalIndentFunc encode indent
type MarshalIndentFunc func(v interface{}, prefix, indent string) ([]byte, error)

// Decoder decoder
type Decoder interface {
	Decode(v interface{}) error
}

// Encoder encoder
type Encoder interface {
	Encode(v interface{}) error
}

var _ JSONDecoder = (*json.Decoder)(nil)

// JSONEncoder interface of *json.Encoder
type JSONEncoder interface {
	Encode(v interface{}) error
	SetEscapeHTML(on bool)
	SetIndent(prefix, indent string)
}

var _ JSONEncoder = (*json.Encoder)(nil)

// XMLDecoder interface of *xml.Decoder
type XMLDecoder interface {
	Decode(v interface{}) error
	DecodeElement(v interface{}, start *xml.StartElement) error
	InputOffset() int64
	RawToken() (xml.Token, error)
	Skip() error
	Token() (xml.Token, error)
}

var _ XMLDecoder = (*xml.Decoder)(nil)

// XMLEncoder interface of *xml.Encoder
type XMLEncoder interface {
	Encode(v interface{}) error
	EncodeElement(v interface{}, start xml.StartElement) error
	EncodeToken(t xml.Token) error
	Flush() error
	Indent(prefix, indent string)
}

var _ XMLEncoder = (*xml.Encoder)(nil)

// CSVReader interface of *csv.Reader
type CSVReader interface {
	Read() (record []string, err error)
	ReadAll() (records [][]string, err error)
}

var _ CSVReader = (*csv.Reader)(nil)

// CSVWriter interface of *csv.Writer
type CSVWriter interface {
	Error() error
	Flush()
	Write(record []string) error
	WriteAll(records [][]string) error
}

var _ CSVWriter = (*csv.Writer)(nil)

// Base32Encoding interface of *base32.Encoding
type Base32Encoding interface {
	Decode(dst, src []byte) (n int, err error)
	DecodeString(s string) ([]byte, error)
	DecodedLen(n int) int
	Encode(dst, src []byte)
	EncodeToString(src []byte) string
	EncodedLen(n int) int
	WithPadding(padding rune) *base32.Encoding
}

var _ Base32Encoding = (*base32.Encoding)(nil)

// Base64Encoding interface of *base64.Encoding
type Base64Encoding interface {
	Decode(dst, src []byte) (n int, err error)
	DecodeString(s string) ([]byte, error)
	DecodedLen(n int) int
	Encode(dst, src []byte)
	EncodeToString(src []byte) string
	EncodedLen(n int) int
	Strict() *base64.Encoding
	WithPadding(padding rune) *base64.Encoding
}

var _ Base64Encoding = (*base64.Encoding)(nil)
