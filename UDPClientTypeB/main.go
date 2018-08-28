package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var defaultDataSize = 1024

func connectToServer(conn *net.UDPConn){
	responseFromServer := make([]byte, 30)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter your command: ")
		time.Sleep(100 * time.Millisecond)
		scanner.Scan()
		userMsg := scanner.Text()//It only reads one line
		_, _ = conn.Write([]byte(userMsg))

		ccc, _ := conn.Read(responseFromServer)
		msg := string(responseFromServer[0:ccc])
		if msg == "Connected" {
			fmt.Printf("Connected to address: %v\n", conn.RemoteAddr())
			fmt.Printf("Connected from address: %v\n", conn.LocalAddr())
			break
		}
	}
}

func readDataFromServer(conn *net.UDPConn){
	 sumOfData  :=0
	buff := make([]byte, defaultDataSize)
	for {
		dataSize, rderr := conn.Read(buff)
		if rderr != nil {
			fmt.Printf("conn.Read() error: %s\n", rderr)
		} else {
			fmt.Printf("Read %d bytes from socket\n", dataSize)
			fmt.Printf("Bytes: %q\n", string(buff[0:dataSize]))
			fmt.Printf("Local address: %v\n", conn.LocalAddr())

			sumOfData += dataSize;
			fmt.Printf("Amount of data: %dKB\n ", sumOfData)
		}
	}
}

func BTypeConnection( portNumber int ){
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{127, 0, 0, 1}, Port: portNumber, Zone: ""})
	if err != nil {
		log.Fatal(err)
	}
	connectToServer(conn)
	readDataFromServer(conn)

	if err = conn.Close(); err != nil {
		log.Fatal(err)
	}
}
func isPortInRange(portNumber int)(bool){
	if portNumber == 0 {
		return false
	}
	if portNumber > 65535 || portNumber <1024 {
		fmt.Println("Invalid port, please try again.")
		return false
	}else{
		return true
	}
}
func run(){
	var portNumber  int= 0

	fmt.Println("Initializing Client B type...")
	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print("Please type the port number of the server you want to connect to(within range 1024-65535): ")

		fmt.Scanf("%d",&portNumber)

		if isPortInRange(portNumber) {
			break
		}
	}
	fmt.Println("To connect to the server type: CONNECT")

	BTypeConnection(portNumber)
}
//TODO DISCONNECT\r\n
//TODO test.go unit test
//For the purpose of the task I've used the localhost address 127.0.0.1

//This loop below demonstrates the amount of B Type clients that you wish to run
//At the moment I've only tested in this UDPClientTypeB project before creating a test unit
//So I could test a large amount of B clients.
//It was used in func run.
//for i := 0; i < 5; i++ {
//
//	go  BTypeConnection(portNumber)
//	time.Sleep(200 * time.Millisecond)
//}
//	  BTypeConnection(portNumber)


func main() {
	run()
}


