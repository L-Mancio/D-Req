package main

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	ConnHost = "192.168.1.51"
	ConnPort = "4444"
	ConnType = "tcp"
)

type reqStructure struct {
	DownloadName string
	ReqAuthor    string
	Priority     int //number 1-5
}

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(ConnType, ConnHost+":"+ConnPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer func(l net.Listener) {
		_ = l.Close()
	}(l)
	fmt.Println("Listening on " + ConnHost + ":" + ConnPort)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	if reqLen > 1024 {
		fmt.Println("Error, message too big")
	}

	readDownloadRequest(buf) //bytes.NewBuffer(buf)
	// Send a response back to person contacting us.
	_, _ = conn.Write([]byte("Request Received"))
	fmt.Println("Request Received, Done")

	// Close the connection when you're done with it.
	_ = conn.Close()
}

func readDownloadRequest(clientReq []byte) { //*bytes.Buffer
	//var reqArray []byte
	file, _ := os.OpenFile("C:\\Users\\lucam\\Desktop\\D-Req\\requests.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	//handles file closure
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()
	rows, _ := csv.NewReader(file).Read()

	if rows == nil {

		firstCsvLine := []string{"Author", "Download Name", "Priority"}
		_ = writer.Write(firstCsvLine)
	}
	//fmt.Printf("%v\n", clientReq)

	dec := gob.NewDecoder(bytes.NewBuffer(clientReq))
	var reqArray []reqStructure
	if err := dec.Decode(&reqArray); err != nil {
		fmt.Println(err)
	}
	for i, singleReq := range reqArray {
		author, downloadName, priority := singleReq.ReqAuthor, singleReq.DownloadName, singleReq.Priority
		lineToAdd := []string{author, downloadName, strconv.FormatInt(int64(priority), 10)}
		fmt.Printf("Request %d: %v, %v, %d\n", i, author, downloadName, priority)
		fmt.Println(lineToAdd)
		_ = writer.Write(lineToAdd)
	}

}
