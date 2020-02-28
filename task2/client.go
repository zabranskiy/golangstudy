package main

import (
	"io"
	"net"
	"os"
)

func main() {
	a, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}

	c, err := net.DialTCP("tcp", nil, a)
	defer c.Close()
	if err != nil {
		panic(err)
	}

	go io.Copy(c, os.Stdin)
	io.Copy(os.Stdout, c)
}