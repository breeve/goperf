package client

import (
	"fmt"
	"time"

	"github.com/panjf2000/gnet/v2"
	"github.com/sirupsen/logrus"
)

func connect(server string, port int, protocol string) {
	// tcpServer, err := net.ResolveTCPAddr(protocol, fmt.Sprintf("%s:%d", server, port))
	// if err != nil {
	// 	logrus.Errorf("client Resolve TCP address fail:%s", err)
	// 	return
	// }
	// conn, err := net.DialTCP(protocol, nil, tcpServer)
	// if err != nil {
	// 	logrus.Errorf("client Dial %s fail:%s", protocol, err)
	// 	return
	// }
	// conn.Write([]byte("this is a client"))

	//cli := &echoClient{}
	evHendler := clientEvents{}
	client, err := gnet.NewClient(
		evHendler,
		gnet.WithTCPNoDelay(gnet.TCPNoDelay),
		gnet.WithLockOSThread(true),
		gnet.WithTicker(true),
	)
	if err != nil {
		logrus.Errorf("new client error:%s", err)
		return
	}
	client.Start()
	defer client.Stop() //nolint:errcheck

	// err = gnet.Run(
	// 	cli,
	// 	fmt.Sprintf("%s://%s:%d", protocol, server, port),
	// 	gnet.WithEdgeTriggeredIO(true),
	// 	gnet.WithLockOSThread(true),
	// 	gnet.WithMulticore(true),
	// 	gnet.WithReusePort(true),
	// 	gnet.WithTicker(true),
	// 	gnet.WithTCPKeepAlive(time.Minute*1),
	// 	gnet.WithLoadBalancing(gnet.RoundRobin),
	// )

	// if err != nil {
	// 	logrus.Errorf("client run error:%s", err)
	// 	return
	// }
	c, err := client.Dial(protocol, fmt.Sprintf("%s:%d", server, port))
	if err != nil {
		logrus.Errorf("client dial error:%s", err)
		return
	}
	defer c.Close()
	err = c.Wake(nil)
	if err != nil {
		logrus.Errorf("client wake error:%s", err)
		return
	}
	_, err = c.Write([]byte("hello"))
	if err != nil {
		logrus.Errorf("client write error:%s", err)
		return
	}
}

type clientEvents struct {
	*gnet.BuiltinEventEngine
	packetLen int
}

type echoClient struct {
	*gnet.BuiltinEventEngine
	eng gnet.Engine
	c   *gnet.Client
}

// OnBoot fires when the engine is ready for accepting connections.
// The parameter engine has information and various utilities.
func (cli *echoClient) OnBoot(eng gnet.Engine) (action gnet.Action) {
	cli.eng = eng
	return gnet.None
}

// OnShutdown fires when the engine is being shut down, it is called right after
// all event-loops and connections are closed.
func (cli *echoClient) OnShutdown(eng gnet.Engine) {
}

// OnOpen fires when a new connection has been opened.
//
// The Conn c has information about the connection such as its local and remote addresses.
// The parameter out is the return value which is going to be sent back to the remote.
// Sending large amounts of data back to the remote in OnOpen is usually not recommended.
func (cli *echoClient) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Printf("local addr:%s", c.LocalAddr())
	fmt.Printf("remote addr:%s", c.RemoteAddr())
	return
}

// OnClose fires when a connection has been closed.
// The parameter err is the last known connection error.
func (cli *echoClient) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	return
}

// OnTraffic fires when a socket receives data from the remote.
//
// Note that the []byte returned from Conn.Peek(int)/Conn.Next(int) is not allowed to be passed to a new goroutine,
// as this []byte will be reused within event-loop after OnTraffic() returns.
// If you have to use this []byte in a new goroutine, you should either make a copy of it or call Conn.Read([]byte)
// to read data into your own []byte, then pass the new []byte to the new goroutine.
func (cli *echoClient) OnTraffic(c gnet.Conn) (action gnet.Action) {
	c.Write([]byte("hello"))
	return
}

// OnTick fires immediately after the engine starts and will fire again
// following the duration specified by the delay return value.
func (cli *echoClient) OnTick() (delay time.Duration, action gnet.Action) {
	return
}
