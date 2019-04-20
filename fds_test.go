// +build linux

package sdialog

import (
	"io"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"testing"
)

func Test_Socket_Activation(t *testing.T) {
	t.Logf("Prepare socket")
	in_r, in_w := socket_pair(t)
	out_r, out_w := socket_pair(t)

	t.Logf("Prepare command")
	cmd := exec.Command("go", "run", "./examples/activation_echo/run.go")
	cmd.ExtraFiles = []*os.File{
		in_r,
		out_w,
	}

	stderr, _ := cmd.StderrPipe()
	env := os.Environ()
	env = append(env, []string{"LISTEN_FDS=2", "NOTIFY_SOCKET=@test"}...)
	cmd.Env = env
	t.Logf("run command")
	err_buf := make([]byte, 65536)
	go io.ReadFull(stderr, err_buf)

	if err := cmd.Start(); err != nil {
		t.Logf(string(err_buf))
		t.Fatalf("run : %v", err)
	}

	t.Logf("write message")
	io.WriteString(in_w, "hello world")
	in_r.Close()
	in_w.Close()

	if err := cmd.Wait(); err != nil {
		t.Logf(string(err_buf))
		t.Fatalf("wait : %v", err)
	}

	buf := make([]byte, 100)
	z, err := io.ReadAtLeast(out_r, buf, 11)
	if err != nil {
		t.Fatalf("read : %v", err)
	}
	if string(buf[0:z]) != "hello world" {
		t.Fatalf("hello world != [%s]", string(buf))

	}
}

func osfilify(fd int, t *testing.T) *os.File {
	return os.NewFile(uintptr(fd), "netconnify_"+strconv.Itoa(fd))
}

func socket_pair(t *testing.T) (*os.File, *os.File) {
	fds, err := syscall.Socketpair(syscall.AF_LOCAL, syscall.SOCK_STREAM, 0)
	if err != nil {
		t.Fatalf("Socketpair: %v", err)
	}

	return osfilify(fds[0], t), osfilify(fds[1], t)
}
