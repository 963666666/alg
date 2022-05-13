package socket

import (
	"bytes"
	"net"
	"time"
)

type ConnConfig struct {
	TcpTransferAfterClose bool

	TcpKeepAlive        bool
	TcpKeepAliveSeconds int

	TcpNoDelay bool

	TcpReadBufferBytes  int
	TcpWriteBufferBytes int
}

type Conn struct {
	connType   string
	conn       net.Conn
	session    map[string]interface{}
	readBuffer bytes.Buffer
	uniqId     int64

	closed bool
}

func (c *Conn) Init(config *ConnConfig, conn *net.TCPConn) {
	if config.TcpTransferAfterClose {
		conn.SetLinger(-1)
	} else {
		conn.SetLinger(0)
	}

	if config.TcpKeepAlive {
		conn.SetKeepAlive(config.TcpKeepAlive)

		if config.TcpKeepAliveSeconds <= 0 {
			config.TcpKeepAliveSeconds = 90
		}
		conn.SetKeepAlivePeriod(time.Duration(config.TcpKeepAliveSeconds) * time.Second)
	}

	conn.SetNoDelay(config.TcpNoDelay)

	if config.TcpReadBufferBytes <= 0 {
		config.TcpReadBufferBytes = 1024
	}

	if config.TcpWriteBufferBytes <= 0 {
		config.TcpWriteBufferBytes = 1024
	}

	c.conn = conn
	c.connType = "socket"
	c.session = make(map[string]interface{})
}
