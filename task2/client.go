package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
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

	snr := bufio.NewScanner(os.Stdin)
	r := bufio.NewReader(c)
	w := bufio.NewWriter(c)

	for snr.Scan() {
		msg := snr.Text() + "\n"
		if len(msg) == 0 {
			break
		}

		_, err = w.WriteString(msg)
		if err == nil {
			err = w.Flush()
		}

		start := time.Now()
		for {
			reply := make([]byte, 1024)
			_, err := r.Read(reply)
			if err != nil {
				panic(err)
			}
			fmt.Print(string(reply))
			elapsed := time.Since(start)
			if elapsed > 2 * time.Second { // лучше не придумал, сервер должен что-то в конце ответа проставлять
				break
			}
		}
	}
}