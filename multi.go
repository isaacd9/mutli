package multi

import (
	"fmt"
	"math/rand"
	"net"
	"os"

	"golang.org/x/sys/unix"
)

const (
	FORK = 0x02
)

func Fork(r Runner) (*os.Process, error) {
	pid_t, _, err := unix.Syscall(FORK, 0, 0, 0)
	if err != 0 {
		return nil, fmt.Errorf("%q", err)
	}
	if int(pid_t) != os.Getpid() {
		return os.FindProcess(int(pid_t))
	} else {
		r.Run()
		fmt.Println()
		os.Exit(0)
	}
	return nil, nil
}

type Process struct {
	todo Runner
	proc *os.Process
}

type Runner interface {
	Run() error
}

func NewProcess(r Runner) *Process {
	return &Process{
		todo: r,
	}
}

func (p *Process) Start() error {
	proc, err := Fork(p.todo)
	if err != nil {
		return fmt.Errorf("Fork(): %v", err)
	}
	p.proc = proc
	return nil
}

func (p *Process) Join() (*os.ProcessState, error) {
	state, err := p.proc.Wait()
	if err != nil {
		return nil, fmt.Errorf("Wait(): %v", err)
	}
	return state, nil
}

type Pipe struct {
	ls   net.Listener
	sock net.Conn
}

func NewPipe() (Pipe, error) {
	ls, err := net.Listen("unix", fmt.Sprintf("/tmp/%d.sock", rand.Uint32()))
	conn, err := net.Dial("unix", fmt.Sprintf("/tmp/%d.sock", rand.Uint32()))
	if err != nil {
		return Pipe{}, fmt.Errorf("Error: %v", err)
	}
	return Pipe{sock: conn, ls: ls}, nil
}

func (p Pipe) Send(b []byte) error {
	_, err := p.sock.Write(b)
	if err != nil {
		return fmt.Errorf("conn.Write: %v", err)
	}
	return nil
}

func (p Pipe) Recv(b []byte) error {
	_, err := p.sock.Read(b)
	if err != nil {
		return fmt.Errorf("conn.Read: %v", err)
	}
	return nil
}

func (p Pipe) Close() {
	p.sock.Close()
}
