package server

import (
	"fmt"

	"github.com/panjf2000/gnet/v2"
	"github.com/sirupsen/logrus"
)

type echoServer struct {
	gnet.BuiltinEventEngine

	eng       gnet.Engine
	addr      string
	multicore bool
}

func (es *echoServer) OnBoot(eng gnet.Engine) gnet.Action {
	es.eng = eng
	return gnet.None
}

func (es *echoServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)
	logrus.Info(string(buf))
	c.Write(buf)
	return gnet.None
}

func listen(address string, port int, protocol string) {
	echo := &echoServer{addr: fmt.Sprintf("%s://%s:%d", protocol, address, port), multicore: true}
	logrus.Fatal(gnet.Run(echo, echo.addr, gnet.WithMulticore(true)))
}
