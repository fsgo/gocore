// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu (duv123+git@baidu.com)
// Date: 2020/12/20

package gocore

import (
	"io"
)

// IOPipeReader interface of *io.PipeReader
type IOPipeReader interface {
	Close() error
	CloseWithError(err error) error
	Read(data []byte) (n int, err error)
}

var _ IOPipeReader = (*io.PipeReader)(nil)

// IOPipeWriter interface of *io.PipeWriter
type IOPipeWriter interface {
	Close() error
	CloseWithError(err error) error
	Write(data []byte) (n int, err error)
}

var _ IOPipeWriter = (*io.PipeWriter)(nil)

// IOSectionReader interface of *io.SectionReader
type IOSectionReader interface {
	Read(p []byte) (n int, err error)
	ReadAt(p []byte, off int64) (n int, err error)
	Seek(offset int64, whence int) (int64, error)
	Size() int64
}

var _ IOSectionReader = (*io.SectionReader)(nil)

// NopWriteCloser close able for writer
func NopWriteCloser(w io.Writer) io.WriteCloser {
	return nopWriteCloser{w}
}

type nopWriteCloser struct {
	io.Writer
}

func (nopWriteCloser) Close() error { return nil }

// Flusher flush
type Flusher interface {
	Flush() error
}
