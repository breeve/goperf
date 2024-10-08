package client

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
)

func connect(server string, port int, protocol string) {
	tcpServer, err := net.ResolveTCPAddr(protocol, fmt.Sprintf("%s:%d", server, port))
	if err != nil {
		logrus.Errorf("client Resolve TCP address fail:%s", err)
		return
	}
	conn, err := net.DialTCP(protocol, nil, tcpServer)
	if err != nil {
		logrus.Errorf("client Dial %s fail:%s", protocol, err)
		return
	}
	conn.Write([]byte("this is a client"))
}
