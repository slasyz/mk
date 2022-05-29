package shell

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/slasyz/mk/src/logger"
)

type Shell struct {
	binary string

	logger *logger.Logger

	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

type Opt func(shell *Shell)

func WithStdin(stdin io.Reader) Opt {
	return func(shell *Shell) {
		shell.stdin = stdin
	}
}
func WithStdout(stdout io.Writer) Opt {
	return func(shell *Shell) {
		shell.stdout = stdout
	}
}
func WithStderr(stderr io.Writer) Opt {
	return func(shell *Shell) {
		shell.stderr = stderr
	}
}
func WithLogger(logger *logger.Logger) Opt {
	return func(shell *Shell) {
		shell.logger = logger
	}
}

func New(binary string, opts ...func(shell *Shell)) *Shell {
	shell := &Shell{
		binary: binary,
		logger: logger.New(nil),
		stdin:  os.Stdin,
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
	for _, opt := range opts {
		opt(shell)
	}
	return shell
}

func (s *Shell) Exec(cmd string, args []string, workdir string) error {
	f, err := ioutil.TempFile("", "mk-script-*.sh")
	if err != nil {
		return fmt.Errorf("error creating temporary script: %w", err)
	}
	defer f.Close()

	err = f.Chmod(0600)
	if err != nil {
		return fmt.Errorf("error changing permissions to temporary script: %w", err)
	}

	_, err = f.WriteString(cmd)
	if err != nil {
		return fmt.Errorf("error writing script to file: %w", err)
	}

	s.logger.Command(cmd, args)

	args = append([]string{f.Name()}, args...)

	command := exec.Command(s.binary, args...)
	command.Dir = workdir
	command.Stdin = s.stdin
	command.Stdout = s.stdout
	command.Stderr = s.stderr

	return command.Run()
}
