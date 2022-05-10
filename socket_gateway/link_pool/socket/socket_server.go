package socket

import (
	"net"
	"strconv"
)

type SocketServerEvent interface {
	OnHandleConn(*Conn)
	OnCloseConn(*Conn)
	OnStopServer()
}

const port = 9001

type SocketServer struct {
	listener *net.TCPListener
	event SocketServerEvent

	waitingFinish chan int
}

func (this *SocketServer) Init(event SocketServerEvent) {
	this.event = event

	this.waitingFinish = make(chan int, 1)
}

func (this *SocketServer) Start() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	this.listener, err = net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	for {
		conn, err := this.listener.AcceptTCP()
		if err != nil {
			<- this.waitingFinish

			return err
		}

		connConfig := new(ConnConfig)
		connConfig.TcpTransferAfterClose = true
		connConfig.TcpKeepAlive = true
		socketConn := new(Conn)
		socketConn.
	}
}