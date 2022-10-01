package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	CommandCd   = "cd"
	CommandPwd  = "pwd"
	CommandEcho = "echo"
	CommandKill = "kill"
	CommandPs   = "ps"
	CommandQuit = "quit"
)

func cd(args []string) ([]byte, error) {
	dir := args[0]
	err := os.Chdir(dir)
	if err != nil {
		return nil, err
	}
	dir, err = os.Getwd()
	if err != nil {
		return nil, err
	}
	return []byte(dir), nil
}

func pwd() ([]byte, error) {
	res, err := os.Getwd()
	return []byte(res), err
}

func echo(args []string) ([]byte, error) {
	return exec.Command(CommandEcho, args...).Output()
}

func kill(pid int) ([]byte, error) {
	process, err := os.FindProcess(pid)
	if err != nil {
		return nil, err
	}
	err = process.Kill()
	if err != nil {
		return nil, err
	}
	return []byte("process killed"), nil
}

func ps() ([]byte, error) {
	return exec.Command(CommandPs).Output()
}

func run(str string) bool {
	var res []byte
	var err error
	args, argsWithSpaces := args(str)
	if !(len(args) > 0) {
		return false
	}
	cmd := args[0]
	switch cmd {
	case CommandCd:
		if len(args) >= 2 {
			res, err = cd(args[1:])
		}
	case CommandPwd:
		res, err = pwd()
	case CommandEcho:
		if len(args) >= 2 {
			res, err = echo(argsWithSpaces[1:])
		}
	case CommandKill:
		if len(args) == 2 {
			pid, err := strconv.Atoi(args[1])
			if err != nil {
				log.Println(err)
			} else {
				res, err = kill(pid)
			}
		}
	case CommandPs:
		res, err = ps()
	case CommandQuit:
		return true
	default:
		log.Println("Invalid arguments")
	}
	if err != nil {
		log.Println(err)
	} else {
		if res != nil {
			fmt.Println(string(res))
		}
	}
	return false
}

func args(str string) ([]string, []string) {
	args := make([]string, 0)
	argsWithSpaces := strings.Split(strings.TrimSpace(str), " ")
	for _, s := range argsWithSpaces {
		if s != "" {
			args = append(args, s)
		}
	}
	return args, argsWithSpaces
}

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	hostname += ":~$ "
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print(hostname)
	for sc.Scan() {
		if isExit := run(sc.Text()); isExit {
			break
		}
		fmt.Print(hostname)
	}
}
