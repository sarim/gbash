package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func isNotSystemdSocket() bool {
	_, err := os.Stat("/var/run/dbus/system_bus_socket")

	return os.IsNotExist(err)
}

func main() {
	for e := isNotSystemdSocket(); e; {
		fmt.Println("Waiting for system socket")
		time.Sleep(100 * time.Millisecond)
		e = isNotSystemdSocket()
	}

	args := []string{"--quiet", "--pty", "--same-dir", "--wait", "--collect", "--service-type=simple", "--uid=" + strconv.Itoa(os.Getuid()), "--property=PAMName=login", "--send-sighup"}

	for _, env := range os.Environ() {
		args = append(args, "--setenv="+env)
	}

	// if _, is := os.LookupEnv("WINDOWS_PATH"); !is {
	// 	args = append(args, "--setenv=WINDOWS_PATH="+os.Getenv("PATH"))
	// }

	shell := "/bin/bash"

	args = append(args, shell)
	args = append(args, "-l")
	args = append(args, os.Args[1:]...)

	c := exec.Command("systemd-run", args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		} else {
			panic(err)
		}
	}
}
