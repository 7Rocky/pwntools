package pwntools

import (
	"fmt"
	"io"
	"net"
	"strings"
)

type ListenConf struct {
	Ip       string
	Protocol string
}

func listen[V number](ip string, port V, protocol string) *Conn {
	p := Progress(fmt.Sprintf("Trying to bind to %s on port %v", ip, port))
	ln, err := net.Listen(protocol, fmt.Sprintf("%s:%v", ip, port))

	if err != nil {
		Error(err.Error())
	}

	defer ln.Close()
	p.Success("Done")
	p = Progress(fmt.Sprintf("Waiting for connections on %s:%v", ip, port))

	c, err := ln.Accept()

	if err != nil {
		Error(err.Error())
	}

	p.Success(fmt.Sprintf("Got connection from %s", c.RemoteAddr().String()))

	hostPort := strings.Split(c.RemoteAddr().String(), ":")
	info = connInfo{host: hostPort[0], port: hostPort[1], isListen: true}
	stdin := io.WriteCloser(c)
	stdout := io.ReadCloser(c)

	return &Conn{stdin: stdin, stdout: stdout, errChan: make(chan error, 1), conn: c}
}

func Listen[V number](port V) *Conn {
	return listen("0.0.0.0", port, "tcp")
}

func ListenWithConf[V number](port V, conf ListenConf) *Conn {
	return listen(conf.Ip, port, conf.Protocol)
}
