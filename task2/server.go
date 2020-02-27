package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connection aborted, err:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	fmt.Println("init connection")
	defer c.Close()
	s := bufio.NewScanner(c)
	for s.Scan() {
		//fmt.Println("scan")
		echo(c, s.Text(), time.Second)
	}
	if err := s.Err(); err != nil {
		fmt.Println("error while scanning, err:", err)
	}
	fmt.Println("close connection")
}

func echo(w io.Writer, in string, d time.Duration) {
	fmt.Println(in)
	_, _ = fmt.Fprintf(w, "\t ((( %s )))", strings.ToUpper(in))
	time.Sleep(d)
	_, _ = fmt.Fprintf(w, "\t (( %s ))", in)
	time.Sleep(d)
	_, _ = fmt.Fprintf(w, "\t ( %s )", strings.ToLower(in))
	time.Sleep(d)
}
