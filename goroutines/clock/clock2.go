package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {

	port := flag.Int("port",8000, "Port number for the server")
	TZ := flag.String("TZ","Asia/Singapore","TimeZone for the timezone clock")

	flag.Parse()

	fmt.Printf("server will run on port %d, for timezone %s \n",*port, *TZ)

	address := fmt.Sprintf("localhost:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		
		locationString := *TZ

		region,err := time.LoadLocation(locationString)
		if err != nil{
			log.Print(err)
		}
		go handleConn(conn,region)
	}

}

func handleConn(c net.Conn, location *time.Location) {
	defer c.Close()

	for {

		currentTime := time.Now().In(location)
		_, err := io.WriteString(c, currentTime.Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
