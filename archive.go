/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/12/20
 */

package gocore

import (
	"archive/tar"
	"archive/zip"
	"io"
	"os"
	"time"
)

// ZipFile interface of *zip.File
type ZipFile interface {
	DataOffset() (offset int64, err error)
	Open() (io.ReadCloser, error)
}

var _ ZipFile = (*zip.File)(nil)

// ZipFileHeader interface of *zip.FileHeader
type ZipFileHeader interface {
	FileInfo() os.FileInfo
	ModTime() time.Time
	Mode() (mode os.FileMode)
	SetModTime(t time.Time)
	SetMode(mode os.FileMode)
}

var _ ZipFileHeader = (*zip.FileHeader)(nil)

// ZipReader interface of *zip.Reader
type ZipReader interface {
	RegisterDecompressor(method uint16, dcomp zip.Decompressor)
}

// ZipWriter interface of *zip.Writer
type ZipWriter interface {
	Close() error
	Create(name string) (io.Writer, error)
	CreateHeader(fh *zip.FileHeader) (io.Writer, error)
	Flush() error
	RegisterCompressor(method uint16, comp zip.Compressor)
	SetComment(comment string) error
	SetOffset(n int64)
}

var _ ZipWriter = (*zip.Writer)(nil)

// TarHeader interface of *tar.Header
type TarHeader interface {
	FileInfo() os.FileInfo
}

var _ TarHeader = (*tar.Header)(nil)

// TarReader interface of *tar.Reader
type TarReader interface {
	Next() (*tar.Header, error)
	Read(b []byte) (int, error)
}

var _ TarReader = (*tar.Reader)(nil)

// TarWriter interface of *tar.Writer
type TarWriter interface {
	Close() error
	Flush() error
	Write(b []byte) (int, error)
	WriteHeader(hdr *tar.Header) error
}

var _ TarWriter = (*tar.Writer)(nil)
