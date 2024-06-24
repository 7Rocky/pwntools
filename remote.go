package pwntools

import (
	"fmt"
	"io"
	"net"
)

type RemoteConf struct {
	Protocol string
}

func remote[V number](host string, port V, protocol string) *Conn {
	info = connInfo{host: host, port: fmt.Sprintf("%v", port), isRemote: true}
	p := Progress(fmt.Sprintf("Opening connection to %s on port %s", info.host, info.port))

	c, err := net.Dial(protocol, fmt.Sprintf("%s:%s", info.host, info.port))
	p.Status(fmt.Sprintf("Trying %s:%s", info.host, info.port))

	if err != nil {
		panic(err)
	}

	stdin := io.WriteCloser(c)
	stdout := io.ReadCloser(c)
	p.Success("Done")

	return &Conn{stdin: stdin, stdout: stdout, errChan: make(chan error, 1)}
}

func Remote[V number](host string, port V) *Conn {
	return remote(host, port, "tcp")
}

func RemoteWithConf[V number](host string, port V, conf RemoteConf) *Conn {
	return remote(host, port, conf.Protocol)
}
