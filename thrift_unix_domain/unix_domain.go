package thrift_unix_domain

import (
	"context"
	"net"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type TTransport struct {
	conn    net.Conn
	addr    net.Addr
	timeout time.Duration
}

// New creates a net.Conn-backed TTransport, given a unix domain file path
//
// Example:
//  trans, err := thrift.New("/tmp/thrift.sock")
func NewTUnixDomain(sockFile string) (*TTransport, error) {
	return NewTUnixDomainTimeout(sockFile, 0)
}

// NewTSocketTimeout creates a net.Conn-backed TTransport, given a unix domain file path
// it also accepts a timeout as a time.Duration
func NewTUnixDomainTimeout(sockFile string, timeout time.Duration) (*TTransport, error) {
	//conn, err := net.DialTimeout(network, address, timeout)
	addr, err := net.ResolveUnixAddr("unix", sockFile)
	if err != nil {
		return nil, err
	}
	return NewTFromAddrTimeout(addr, timeout), nil
}

// Creates a TUnixDomain from a net.Addr
func NewTFromAddrTimeout(addr net.Addr, timeout time.Duration) *TTransport {
	return &TTransport{addr: addr, timeout: timeout}
}

// Sets the socket timeout
func (p *TTransport) SetTimeout(timeout time.Duration) error {
	p.timeout = timeout
	return nil
}

func (p *TTransport) pushDeadline(read, write bool) {
	var t time.Time
	if p.timeout > 0 {
		t = time.Now().Add(time.Duration(p.timeout))
	}
	if read && write {
		p.conn.SetDeadline(t)
	} else if read {
		p.conn.SetReadDeadline(t)
	} else if write {
		p.conn.SetWriteDeadline(t)
	}
}

// Connects the socket, creating a new socket object if necessary.
func (p *TTransport) Open() error {
	if p.IsOpen() {
		return thrift.NewTTransportException(thrift.ALREADY_OPEN, "Socket already connected.")
	}
	if p.addr == nil {
		return thrift.NewTTransportException(thrift.NOT_OPEN, "Cannot open nil address.")
	}
	if len(p.addr.Network()) == 0 {
		return thrift.NewTTransportException(thrift.NOT_OPEN, "Cannot open bad network name.")
	}
	if len(p.addr.String()) == 0 {
		return thrift.NewTTransportException(thrift.NOT_OPEN, "Cannot open bad address.")
	}
	var err error
	if p.conn, err = net.DialTimeout(p.addr.Network(), p.addr.String(), p.timeout); err != nil {
		return thrift.NewTTransportException(thrift.NOT_OPEN, err.Error())
	}
	return nil
}

// Retreive the underlying net.Conn
func (p *TTransport) Conn() net.Conn {
	return p.conn
}

// Returns true if the connection is open
func (p *TTransport) IsOpen() bool {
	if p.conn == nil {
		return false
	}
	return true
}

// Closes the socket.
func (p *TTransport) Close() error {
	// Close the socket
	if p.conn != nil {
		err := p.conn.Close()
		if err != nil {
			return err
		}
		p.conn = nil
	}
	return nil
}

func (p *TTransport) Read(buf []byte) (int, error) {
	if !p.IsOpen() {
		return 0, thrift.NewTTransportException(thrift.NOT_OPEN, "Connection not open")
	}
	p.pushDeadline(true, false)
	n, err := p.conn.Read(buf)
	return n, thrift.NewTTransportExceptionFromError(err)
}

func (p *TTransport) Write(buf []byte) (int, error) {
	if !p.IsOpen() {
		return 0, thrift.NewTTransportException(thrift.NOT_OPEN, "Connection not open")
	}
	p.pushDeadline(false, true)
	return p.conn.Write(buf)
}

func (p TTransport) Peek() bool {
	return p.IsOpen()
}

func (p *TTransport) Flush(ctx context.Context) error {
	return nil
}

func (p *TTransport) Interrupt() error {
	if !p.IsOpen() {
		return nil
	}
	return p.conn.Close()
}

func (p *TTransport) RemainingBytes() (num_bytes uint64) {
	const maxSize = ^uint64(0)
	return maxSize // the thruth is, we just don't know unless framed is used
}