package pwntools

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func Listen[V Number](port V) *Conn {
	p := Progress(fmt.Sprintf("Trying to bind to 0.0.0.0 on port %v", port))
	ln, _ := net.Listen("tcp", fmt.Sprintf(":%v", port))

	p.Status("Trying ::")
	p.Success("Done")

	p = Progress(fmt.Sprintf("Waiting for connections on 0.0.0.0:%v", port))

	c, _ := ln.Accept()
	p.Success(fmt.Sprintf("Got connection from %s", c.RemoteAddr().String()))
	hostPort := strings.Split(c.RemoteAddr().String(), ":")
	info = connInfo{host: hostPort[0], port: hostPort[1], isLocal: true}
	stdin := io.WriteCloser(c)
	stdout := io.ReadCloser(c)

	return &Conn{stdin: stdin, stdout: stdout}
}
