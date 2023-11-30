package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"text/tabwriter"
	"time"

)

func main() {
	// Define flags for the clock servers.
	newYork := flag.String("NewYork", "", "New York clock server address")
	london := flag.String("London", "", "London clock server address")
	tokyo := flag.String("Tokyo", "", "Tokyo clock server address")

	flag.Parse()

	if *newYork == "" || *london == "" || *tokyo == "" {
		fmt.Println("Please provide addresses for all clock servers.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Connect to clock servers.
	newYorkConn := connectToClockServer(*newYork)
	londonConn := connectToClockServer(*london)
	tokyoConn := connectToClockServer(*tokyo)

	defer newYorkConn.Close()
	defer londonConn.Close()
	defer tokyoConn.Close()

	// Read and display times from clock servers.
	displayClocks(newYorkConn, londonConn, tokyoConn)

}

func connectToClockServer(address string) net.Conn {
	Conn,err := net.Dial("tcp",address)
	if err!= nil{
		log.Fatal(err)
	}
	return Conn
}

func displayClocks(newYorkConn, londonConn, tokyoConn net.Conn) {
	// Use tabwriter to format the output as a table.
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	for {
		// Read times from each clock server.
		newYorkTime := readTime(newYorkConn)
		londonTime := readTime(londonConn)
		tokyoTime := readTime(tokyoConn)

		// Display the times in a table.
		fmt.Fprintf(w, "New York:\t%s\tLondon:\t%s\tTokyo:\t%s\t\n",
			newYorkTime.Format("15:04:05"), londonTime.Format("15:04:05"), tokyoTime.Format("15:04:05"))

		// Flush the buffer to ensure the output is printed.
		w.Flush()

		// Wait for a short interval before reading the times again.
		time.Sleep(1 * time.Second)
	}
}

func readTime(conn net.Conn) time.Time {
	var timeStr string

	_, err := fmt.Fscanf(conn, "%s\n", &timeStr)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the time string.
	t, err := time.Parse("15:04:05", timeStr)
	if err != nil {
		log.Fatal(err)
	}

	return t
}