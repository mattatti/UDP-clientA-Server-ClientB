package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func writeDataToServer(conn *net.UDPConn){
	fmt.Printf("Connected: %T, %v\n", conn, conn)
	fmt.Printf("Local address: %v\n", conn.LocalAddr())
	fmt.Printf("Remote address: %v\n", conn.RemoteAddr())

	buff := make([]byte, 1024)

	var counter=1
	var packetNumber=1
	for {
		//Counter used just so I could see the difference with data in each packet
		buff[counter]=1
		counter++
		if counter == 1000 {
			counter = 0
		}
		cc, wrerr := conn.Write(buff)

		if wrerr != nil {
			fmt.Printf("conn.Write() error: %s\n", wrerr)
		} else {
			fmt.Printf("packet #%d, Wrote %d bytes to socket\n",packetNumber, cc)
			packetNumber++
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func ATypeConnection(portNumber int){
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP:[]byte{127,0,0,1},Port:portNumber,Zone:""})
	if err != nil {
		log.Fatal(err)
	}
	writeDataToServer(conn)

	if err = conn.Close(); err != nil {
		log.Fatal(err)
	}
}
func isPortInRange(portNumber int)(bool){
	if portNumber > 65535 || portNumber <1024 {
		fmt.Println("Invalid port, please try again.")
		return false
	}else{
		return true
	}
}
func run(){
	portNumber := 0

	fmt.Println("Initializing Client A type...")
	for {
		fmt.Print("Please type the port number of the server you want to connect to(within range 1024-65535): ")
		fmt.Scanf("%d\n", &portNumber)
		if isPortInRange(portNumber) {
			break
		}
	}
	ATypeConnection(portNumber)
}

//For the purpose of the task I've used the localhost address 127.0.0.1
func main() {
	run()
}
