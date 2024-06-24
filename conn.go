package pwntools

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"syscall"
	"time"

	"os/signal"
)

type Conn struct {
	stdin    io.WriteCloser
	stdout   io.ReadCloser
	isClosed bool
	errChan  chan error
}

type connInfo struct {
	command   string
	pid       int
	host      string
	port      string
	isListen  bool
	isProcess bool
	isRemote  bool
}

var info connInfo

func (conn *Conn) writeInteractive(prompt string) {
	for {
		b := conn.Recv()

		if len(b) == 0 {
			conn.errChan <- fmt.Errorf("EOF")
			return
		}

		fmt.Print("\r")
		os.Stdout.Write(b)
		fmt.Print(prompt)
	}
}

func (conn *Conn) readInteractive(reader *bufio.Reader, finish chan struct{}, prompt string) {
	for {
		select {
		case <-finish:
			return
		default:
			s, _ := reader.ReadString('\n')

			if len(s) > 0 {
				conn.Send([]byte(s))
				fmt.Print(prompt)
			}
		}
	}
}

func (conn *Conn) Interactive(prompt ...string) {
	Info("Switching to interactive mode")

	if len(prompt) == 0 {
		prompt = append(prompt, fmt.Sprintf("%s$%s ", red, reset))
	}

	fmt.Print(prompt[0])

	finish := make(chan struct{}, 1)
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-c
		conn.errChan <- fmt.Errorf("Control-C")
	}()

	syscall.SetNonblock(int(os.Stdin.Fd()), true)
	reader := bufio.NewReader(os.Stdin)

	go conn.writeInteractive(prompt[0])
	go conn.readInteractive(reader, finish, prompt[0])

	for {
		err := <-conn.errChan

		if err != nil {
			if err.Error() == "EOF" {
				Info("Got EOF while reading in interactive")
			}

			if err.Error() == "Control-C" {
				Info("Interrupted")
			}

			os.Stdin.SetDeadline(time.Now())
			finish <- struct{}{}
			time.Sleep(10 * time.Millisecond)
			break
		}
	}

	syscall.SetNonblock(int(os.Stdin.Fd()), false)

	close(finish)
	conn.Close()
}

func (conn *Conn) Close() {
	if !conn.isClosed {
		conn.isClosed = true

		conn.stdin.Close()
		conn.stdout.Close()

		if info.isProcess {
			Info("Stopped process '%s' (pid %d)\n", info.command, info.pid)
		} else if info.isRemote {
			Info("Closed connection to %s port %s\n", info.host, info.port)
		} else if info.isListen {
			Info("Closed connection to %s port %s\n", info.host, info.port)
		}
	}
}

func (conn *Conn) Recv(n ...int) []byte {
	max := 4096

	if len(n) == 1 {
		max = n[0]
	} else if len(n) > 1 {
		panic("multiple n")
	}

	buf := make([]byte, max)
	read, err := conn.stdout.Read(buf)

	if err != nil {
		return []byte{}
	}

	return buf[:read]
}

func (conn *Conn) RecvN(n int) []byte {
	buf := make([]byte, n)
	read, err := conn.stdout.Read(buf)

	if err != nil {
		Error(err.Error())
	}

	if read == n {
		return buf
	}

	return []byte{}
}

func (conn *Conn) RecvUntil(pattern []byte, drop ...bool) []byte {
	var recv []byte
	b := make([]byte, 1)

	for !bytes.HasSuffix(recv, pattern) {
		n, err := conn.stdout.Read(b)

		if err != nil {
			panic(err)
		}

		if n == 1 {
			recv = append(recv, b[0])
		}
	}

	if len(drop) == 1 && drop[0] {
		return bytes.ReplaceAll(recv, pattern, []byte(""))
	}

	return recv
}

func (conn *Conn) RecvLine() []byte {
	return conn.RecvUntil([]byte{'\n'})
}

func (conn *Conn) RecvLineContains(pattern []byte) []byte {
	var recv []byte

	for {
		recv = append(recv, conn.RecvLine()...)

		if bytes.Contains(recv, pattern) {
			return recv
		}
	}
}

func (conn *Conn) Send(data []byte) int {
	n, err := conn.stdin.Write(data)

	if err != nil {
		Error(err.Error())
	}

	return n
}

func (conn *Conn) SendLine(data []byte) int {
	return conn.Send(append(data, '\n'))
}

func (conn *Conn) SendLineAfter(pattern, data []byte) []byte {
	recv := conn.RecvUntil(pattern)
	conn.SendLine(data)
	return recv
}

func (conn *Conn) SendAfter(pattern, data []byte) []byte {
	recv := conn.RecvUntil(pattern)
	conn.Send(data)
	return recv
}

func (conn *Conn) RecvS(max ...int) string {
	return string(conn.Recv(max...))
}

func (conn *Conn) RecvNS(n int) string {
	return string(conn.RecvN(n))
}

func (conn *Conn) RecvUntilS(pattern []byte, drop ...bool) string {
	return string(conn.RecvUntil(pattern, drop...))
}

func (conn *Conn) RecvLineS() string {
	return string(conn.RecvLine())
}

func (conn *Conn) RecvLineContainsS(pattern []byte) string {
	return string(conn.RecvLineContains(pattern))
}
