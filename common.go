/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/12/20
 */

package gocore

// Timeout Timeout() bool
type Timeout interface {
	Timeout() bool
}

// IsTimeout check is timeout
func IsTimeout(err interface{}) bool {
	if e, ok := err.(Timeout); ok {
		return e.Timeout()
	}
	return false
}

// Temporary Temporary() bool
type Temporary interface {
	Temporary() bool
}

// IsTemporary check is Temporary
func IsTemporary(err interface{}) bool {
	if e, ok := err.(Temporary); ok {
		return e.Temporary()
	}
	return false
}

// Reset reset
type Reset interface {
	Reset()
}

// ResetErr reset and return error
type ResetErr interface {
	Reset() error
}

// Truncate truncate
type Truncate interface {
	Truncate(n int)
}
