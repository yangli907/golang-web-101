package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8088")

	if err != nil {
		log.Panic(err)
	} 

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	i := 0

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		// parsing logic to extract key info from http request
		if i == 0 {
			method := strings.Fields(ln)[0] // http method
			uri := strings.Fields(ln)[1]

			fmt.Println("METHOD is:", method)
			fmt.Println("URI is:", uri)
		}

		if ln == "" {
			break
		}
		i++
	}

	writeResponse(conn)	

	defer conn.Close()
}

func writeResponse(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"> <head><meta charet="UTF-8"></head><body><string> hello there</strong> </body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}