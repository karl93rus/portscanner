package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

func main() {
	host := os.Args[1]
	// port := os.Args[2]

	for i := 1; i < 65535; i++ {
		hostPort := host + ":" + strconv.Itoa(i)
		fmt.Println(hostPort)
		conn, err := net.Dial("tcp", hostPort)
		if err != nil {
			continue
		} else {
			fmt.Println("TCP port open: " + strconv.Itoa(i))
			continue
		}
		defer conn.Close()
		io.Copy(os.Stdin, conn)
	}
}
