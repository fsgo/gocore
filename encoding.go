/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/12/20
 */

package gocore

// Marshal encode
type Marshal func(v interface{}) ([]byte, error)

// Unmarshal decode
type Unmarshal func(data []byte, v interface{}) error

// MarshalIndent encode indent
type MarshalIndent func(v interface{}, prefix, indent string) ([]byte, error)
