package pwntools

import (
	"fmt"
	"os/exec"
)

type ProcessConf struct {
	Env       []string
	IgnoreEnv bool
	Cwd       string
}

func process(cmd *exec.Cmd) *Conn {
	p := Progress(fmt.Sprintf("Starting local process '%s'", cmd.Path))

	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	conn := Conn{stdin: stdin, stdout: stdout}

	cmd.Start()
	go wait(cmd, &conn)

	info = connInfo{command: cmd.Path, pid: cmd.Process.Pid, isProcess: true}
	p.Success(fmt.Sprintf("pid %d", info.pid))

	return &conn
}

func Process(command string, args ...string) *Conn {
	return process(exec.Command(command, args...))
}

func ProcessWithConf(argv []string, conf ProcessConf) *Conn {
	if len(argv) == 0 {
		Error("Empty argv")
	}

	cmd := exec.Command(argv[0], argv[1:]...)

	if len(conf.Env) != 0 {
		if conf.IgnoreEnv {
			cmd.Env = conf.Env
		} else {
			cmd.Env = append(cmd.Env, conf.Env...)
		}
	}

	if conf.Cwd != "" {
		cmd.Dir = conf.Cwd
	}

	return process(cmd)
}

func wait(cmd *exec.Cmd, conn *Conn) {
	if err := cmd.Wait(); err != nil {
		Warning("Command finished with error: %v", err)
		conn.toClose = true
	}
}
