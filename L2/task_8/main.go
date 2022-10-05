package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	CommandCd   = "cd"
	CommandLs   = "ls"
	CommandPwd  = "pwd"
	CommandEcho = "echo"
	CommandKill = "kill"
	CommandPs   = "ps"
	CommandQuit = "quit"
)

func cd(args []string) (string, error) {
	dir := args[0]
	err := os.Chdir(dir)
	if err != nil {
		return "", err
	}
	dir, err = os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

func ls() (string, error) {
	files, err := os.ReadDir(".")
	if err != nil {
		return "", err
	}
	s := strings.Builder{}
	for i, f := range files {
		s.WriteString(f.Name())
		if i < len(files)-1 {
			s.WriteString("\n")
		}
	}
	return s.String(), nil
}

func pwd() (string, error) {
	return os.Getwd()
}

func echo(args []string) (string, error) {
	res, err := exec.Command(CommandEcho, args...).Output()
	return string(res), err
}

func kill(pid int) (string, error) {
	process, err := os.FindProcess(pid)
	if err != nil {
		return "", err
	}
	err = process.Kill()
	if err != nil {
		return "", err
	}
	return "process killed", nil
}

func ps() (string, error) {
	res, err := exec.Command(CommandPs).Output()
	return string(res), err
}

func run(str string) bool {
	var res string
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
	case CommandLs:
		res, err = ls()
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
		if len(res) > 0 {
			fmt.Println(res)
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
	green := color.New(color.FgGreen).SprintfFunc()
	hostname += ":~$ "
	sc := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s", green(hostname))
	for sc.Scan() {
		if isExit := run(sc.Text()); isExit {
			break
		}
		fmt.Printf("%s", green(hostname))
	}
}
