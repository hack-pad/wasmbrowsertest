package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func main() {
	err := run(os.Args[1:])
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		os.Exit(exitErr.ExitCode())
	}
	if err != nil {
		os.Exit(1)
	}
}

func run(args []string) error {
	env := os.Environ()
	var smallEnv []string
	for _, keyValue := range env {
		tokens := strings.SplitN(keyValue, "=", 2)
		if allowEnvName(tokens[0]) {
			smallEnv = append(smallEnv, keyValue)
		}
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = smallEnv
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func allowEnvName(name string) bool {
	switch {
	case
		strings.HasPrefix(name, "GITHUB_"),
		strings.HasPrefix(name, "JAVA_"):
		return false
	default:
		return true
	}
}
