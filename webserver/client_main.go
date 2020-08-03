package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8088")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Fprintln(conn, "request sent!")
}