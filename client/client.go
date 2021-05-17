package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)
type reqStructure struct{
	DownloadName string
	ReqAuthor string
	Priority int //number 1-5
}
func main() {
	fmt.Println("Enter your name:")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	var reqArray []reqStructure

	if name!= ""{
		for  {
			fmt.Println("Enter Download name, be specific")
			reader := bufio.NewReader(os.Stdin)
			dwnldName, _ := reader.ReadString('\n')
			dwnldName = strings.TrimSuffix(dwnldName, "\n")
			fmt.Println("Enter priority a number from 1-5, 5 being highest priority")
			prio, _ := reader.ReadString('\n')
			prio = strings.TrimSuffix(prio, "\n")
			prioInt, _ := strconv.ParseInt(prio, 10, 64)
			fmt.Println("are you done? y/n")
			var resp string
			_, _ = fmt.Scanln(&resp) // used scanln since reader gives problems on CMD

			reqArray = append(reqArray, reqStructure{
				DownloadName: dwnldName,
				ReqAuthor:    name,
				Priority: int(prioInt),
			})
			if strings.ToLower(resp) == "y"{
				break
			}
		}
	}
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(reqArray); err != nil {
		fmt.Println(err)
	}

	requestInByteFormat := buf.Bytes()

	servAddr := "IP:4444"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write(requestInByteFormat)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("Sending to server ", servAddr)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	//println("reply from server=", string(reply))

	_ = conn.Close()
}
