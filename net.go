// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu (duv123+git@baidu.com)
// Date: 2020/12/20

package gocore

import (
	"context"
	"io"
	"net"
	"os"
	"syscall"
	"time"
)

// NetIPConn interface of *net.IPConn
type NetIPConn interface {
	Close() error
	File() (f *os.File, err error)
	LocalAddr() net.Addr
	Read(b []byte) (int, error)
	ReadFrom(b []byte) (int, net.Addr, error)
	ReadFromIP(b []byte) (int, *net.IPAddr, error)
	ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *net.IPAddr, err error)
	RemoteAddr() net.Addr
	SetDeadline(t time.Time) error
	SetReadBuffer(bytes int) error
	SetReadDeadline(t time.Time) error
	SetWriteBuffer(bytes int) error
	SetWriteDeadline(t time.Time) error
	SyscallConn() (syscall.RawConn, error)
	Write(b []byte) (int, error)
	WriteMsgIP(b, oob []byte, addr *net.IPAddr) (n, oobn int, err error)
	WriteTo(b []byte, addr net.Addr) (int, error)
	WriteToIP(b []byte, addr *net.IPAddr) (int, error)
}

var _ NetIPConn = (*net.IPConn)(nil)

// NetListenConfig interface of *net.ListenConfig
type NetListenConfig interface {
	Listen(ctx context.Context, network, address string) (net.Listener, error)
	ListenPacket(ctx context.Context, network, address string) (net.PacketConn, error)
}

var _ NetListenConfig = (*net.ListenConfig)(nil)

// NetOpError interface of *net.OpError
type NetOpError interface {
	Error() string
	Temporary() bool
	Timeout() bool
	Unwrap() error
}

var _ NetOpError = (*net.OpError)(nil)

// NetTCPConn interface of *net.TCPConn
type NetTCPConn interface {
	Close() error
	CloseRead() error
	CloseWrite() error
	File() (f *os.File, err error)
	LocalAddr() net.Addr
	Read(b []byte) (int, error)
	ReadFrom(r io.Reader) (int64, error)
	RemoteAddr() net.Addr
	SetDeadline(t time.Time) error
	SetKeepAlive(keepalive bool) error
	SetKeepAlivePeriod(d time.Duration) error
	SetLinger(sec int) error
	SetNoDelay(noDelay bool) error
	SetReadBuffer(bytes int) error
	SetReadDeadline(t time.Time) error
	SetWriteBuffer(bytes int) error
	SetWriteDeadline(t time.Time) error
	SyscallConn() (syscall.RawConn, error)
	Write(b []byte) (int, error)
}

var _ NetTCPConn = (*net.TCPConn)(nil)

// NetTCPListener interface of *net.TCPListener
type NetTCPListener interface {
	Accept() (net.Conn, error)
	AcceptTCP() (*net.TCPConn, error)
	Addr() net.Addr
	Close() error
	File() (f *os.File, err error)
	SetDeadline(t time.Time) error
	SyscallConn() (syscall.RawConn, error)
}

var _ NetTCPListener = (*net.TCPListener)(nil)

// NetUDPConn interface of *net.UDPConn
type NetUDPConn interface {
	Close() error
	File() (f *os.File, err error)
	LocalAddr() net.Addr
	Read(b []byte) (int, error)
	ReadFrom(b []byte) (int, net.Addr, error)
	ReadFromUDP(b []byte) (int, *net.UDPAddr, error)
	ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *net.UDPAddr, err error)
	RemoteAddr() net.Addr
	SetDeadline(t time.Time) error
	SetReadBuffer(bytes int) error
	SetReadDeadline(t time.Time) error
	SetWriteBuffer(bytes int) error
	SetWriteDeadline(t time.Time) error
	SyscallConn() (syscall.RawConn, error)
	Write(b []byte) (int, error)
	WriteMsgUDP(b, oob []byte, addr *net.UDPAddr) (n, oobn int, err error)
	WriteTo(b []byte, addr net.Addr) (int, error)
	WriteToUDP(b []byte, addr *net.UDPAddr) (int, error)
}

var _ NetUDPConn = (*net.UDPConn)(nil)

// NetUnixConn interface of *net.UnixConn
type NetUnixConn interface {
	Close() error
	CloseRead() error
	CloseWrite() error
	File() (f *os.File, err error)
	LocalAddr() net.Addr
	Read(b []byte) (int, error)
	ReadFrom(b []byte) (int, net.Addr, error)
	ReadFromUnix(b []byte) (int, *net.UnixAddr, error)
	ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *net.UnixAddr, err error)
	RemoteAddr() net.Addr
	SetDeadline(t time.Time) error
	SetReadBuffer(bytes int) error
	SetReadDeadline(t time.Time) error
	SetWriteBuffer(bytes int) error
	SetWriteDeadline(t time.Time) error
	SyscallConn() (syscall.RawConn, error)
	Write(b []byte) (int, error)
	WriteMsgUnix(b, oob []byte, addr *net.UnixAddr) (n, oobn int, err error)
	WriteTo(b []byte, addr net.Addr) (int, error)
	WriteToUnix(b []byte, addr *net.UnixAddr) (int, error)
}

var _ NetUnixConn = (*net.UnixConn)(nil)

// NetUnixListener interface of *net.UnixListener
type NetUnixListener interface {
	Accept() (net.Conn, error)
	AcceptUnix() (*net.UnixConn, error)
	Addr() net.Addr
	Close() error
	File() (f *os.File, err error)
	SetDeadline(t time.Time) error
	SetUnlinkOnClose(unlink bool)
	SyscallConn() (syscall.RawConn, error)
}

var _ NetUnixListener = (*net.UnixListener)(nil)

var _ NetResolver = (*net.Resolver)(nil)

// NetAccept accept net conn
type NetAccept interface {
	Accept() (net.Conn, error)
}

// Serve serve listener
type Serve interface {
	Serve(l net.Listener) error
}

// ServeTLS ServeTLS
type ServeTLS interface {
	ServeTLS(l net.Listener, certFile, keyFile string) error
}

// ShutdownContext shutdown
type ShutdownContext interface {
	Shutdown(ctx context.Context) error
}
