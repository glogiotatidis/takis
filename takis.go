package main

import "net"
import "time"
import "os"
import "fmt"

func main() {
	host := os.Getenv("CHECK_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("CHECK_PORT")
	if port == "" {
		fmt.Printf("Need port")
		return
	}

	for {
		_, err := net.Dial("tcp", host + ":" + port)
		if err == nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
