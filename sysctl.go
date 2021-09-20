package sysctl

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strings"
)

const (
	sysctlDir = "/proc/sys/"
)

var invalidKeyError = errors.New("could not find the given key")

func linuxGet(name string) (string, error) {
	path := sysctlDir + strings.Replace(name, ".", "/", -1)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", invalidKeyError
	}
	return strings.TrimSpace(string(data)), nil
}

func freebsdGet(name string) (string, error) {
	var stdout bytes.Buffer

	cmd := exec.Command("sysctl", "-n", name)
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return strings.TrimRight(stdout.String(), "\n"), nil
}

func Get(name string) (string, error) {
	switch runtime.GOOS {
	case "linux":
		return linuxGet(name)
	case "freebsd":
		return freebsdGet(name)
	default:
		return "", fmt.Errorf("sysctl: This runtime is not supported: %v", runtime.GOOS)
	}
}

func linuxSet(name string, value string) error {
	path := sysctlDir + strings.Replace(name, ".", "/", -1)
	err := ioutil.WriteFile(path, []byte(value), 0644)
	return err
}

func freebsdSet(name string, value string) error {
	return exec.Command("sysctl", fmt.Sprintf("%s=%s", name, value)).Run()
}

func Set(name string, value string) error {
	switch runtime.GOOS {
	case "linux":
		return linuxSet(name, value)
	case "freebsd":
		return freebsdSet(name, value)
	default:
		return fmt.Errorf("sysctl: This runtime is not supported: %v", runtime.GOOS)
	}
}
