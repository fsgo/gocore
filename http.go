// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu (duv123+git@baidu.com)
// Date: 2020/12/20

package gocore

import (
	"context"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
)

// HTTPClient interface of *http.Client
type HTTPClient interface {
	CloseIdleConnections()
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (resp *http.Response, err error)
	Head(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}

var _ HTTPClient = (*http.Client)(nil)

// HTTPRequest interface of *http.Request
type HTTPRequest interface {
	AddCookie(c *http.Cookie)
	BasicAuth() (username, password string, ok bool)
	Clone(ctx context.Context) *http.Request
	Context() context.Context
	Cookie(name string) (*http.Cookie, error)
	Cookies() []*http.Cookie
	FormFile(key string) (multipart.File, *multipart.FileHeader, error)
	FormValue(key string) string
	MultipartReader() (*multipart.Reader, error)
	ParseForm() error
	ParseMultipartForm(maxMemory int64) error
	PostFormValue(key string) string
	ProtoAtLeast(major, minor int) bool
	Referer() string
	SetBasicAuth(username, password string)
	UserAgent() string
	WithContext(ctx context.Context) *http.Request
	Write(w io.Writer) error
	WriteProxy(w io.Writer) error
}

var _ HTTPRequest = (*http.Request)(nil)

// HTTPResponse interface of *http.Response
type HTTPResponse interface {
	Cookies() []*http.Cookie
	Location() (*url.URL, error)
	ProtoAtLeast(major, minor int) bool
	Write(w io.Writer) error
}

var _ HTTPResponse = (*http.Response)(nil)

// HTTPTransport interface of *http.Transport
type HTTPTransport interface {
	CancelRequest(req *http.Request)
	Clone() *http.Transport
	CloseIdleConnections()
	RegisterProtocol(scheme string, rt http.RoundTripper)
	RoundTrip(req *http.Request) (*http.Response, error)
}

var _ HTTPTransport = (*http.Transport)(nil)

// HTTPServer interface of *http.Server
type HTTPServer interface {
	Close() error
	ListenAndServe() error
	ListenAndServeTLS(certFile, keyFile string) error
	RegisterOnShutdown(f func())
	Serve(l net.Listener) error
	ServeTLS(l net.Listener, certFile, keyFile string) error
	SetKeepAlivesEnabled(v bool)
	Shutdown(ctx context.Context) error
}

var _ HTTPServer = (*http.Server)(nil)

// HTTPServeMux  interface of *http.ServeMux
type HTTPServeMux interface {
	Handle(pattern string, handler http.Handler)
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	Handler(r *http.Request) (h http.Handler, pattern string)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

var _ HTTPServeMux = (*http.ServeMux)(nil)

// RegisterOnShutdown register func OnShutdown
type RegisterOnShutdown interface {
	RegisterOnShutdown(f func())
}
