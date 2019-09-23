package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

type Host struct {
	host    string
	minPort int
	curPort int
	maxPort int
	lock    sync.Mutex
}

func (h *Host) scan() {
	fmt.Printf("In method: %v\n", h)
	for h.curPort < h.maxPort {
		h.lock.Lock()
		hostPort := h.host + ":" + strconv.Itoa(h.curPort)
		fmt.Println("In method. Host: ", hostPort)
		conn, err := net.DialTimeout("tcp", hostPort, time.Millisecond*500)
		if err != nil {
			continue
		} else {
			fmt.Println("TCP port open: " + strconv.Itoa(h.curPort))
			continue
		}
		conn.Close()
		h.curPort++
		h.lock.Unlock()
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	host := os.Args[1]

	h := &Host{
		host:    host,
		minPort: 1,
		curPort: 1,
		maxPort: 10000,
	}

	fmt.Printf("In main. Host: %v\n", h.host)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go h.scan()
	}
	wg.Wait()

	// scanPorts(host)
}

func scanPorts(host string) {
	for i := 77; i < 10000; i++ {
		hostPort := host + ":" + strconv.Itoa(i)
		fmt.Println(hostPort)
		conn, err := net.DialTimeout("tcp", hostPort, time.Millisecond*500)
		if err != nil {
			continue
		} else {
			fmt.Println("TCP port open: " + strconv.Itoa(i))
			continue
		}
		conn.Close()
	}
}
