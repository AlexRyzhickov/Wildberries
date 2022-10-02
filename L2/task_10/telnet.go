package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.String("timeout", "10s", "timeout")
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("Invalid args")
	}
	host := args[0]
	port := args[1]
	duration, err := time.ParseDuration(*timeout)
	if err != nil {
		log.Fatal(err)
	}
	connection, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), duration)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	log.Println("Connection established")

	errChan := make(chan error)
	in := make(chan string)
	out := make(chan string)

	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
		<-shutdown
		errChan <- errors.New("aborted by system interrupt")
	}()

	go read(connection, in, errChan)
	go read(os.Stdin, out, errChan)

Loop:
	for {
		select {
		case in := <-in:
			fmt.Print(in)
		case out := <-out:
			_, err := connection.Write([]byte(out))
			if err != nil {
				errChan <- err
			}
		case err := <-errChan:
			log.Println(err)
			break Loop
		}
	}
}

func read(in io.Reader, channel chan string, errChan chan error) {
	r := bufio.NewReader(in)
	for {
		str, err := r.ReadString('\n')
		if err != nil {
			errChan <- err
		}
		channel <- str
	}
}
