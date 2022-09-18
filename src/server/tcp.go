package server

import (
	"fmt"
	"net"
)

type AbcTcpServer struct {
	listener net.Listener
}

func (tcpss *AbcTcpServer) ListenConn(handler func(conn *TunnelConn)) {
	for {
		conn, err := tcpss.listener.Accept()
		if err != nil {
			continue
		}
		wrapConn := &TunnelConn{Tuntype: "tcp", tcpSocket: conn}
		handler(wrapConn)
	}
}

func NewAbcTcpServer(host string, port int) (svc *AbcTcpServer, err error) {
	address := fmt.Sprintf("%s:%d", host, port)
	tcpserver, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	return &AbcTcpServer{listener: tcpserver}, nil
}