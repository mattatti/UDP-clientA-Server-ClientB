package main

import (
	"fmt"
	"log"
	"net"
)

var packet = make([]byte, 1024)
var isAllowedClientAFlag=false
var dataSizeForBClient =0

func listen( conn2 *net.UDPConn,err2 error, quit chan struct{} ) {
	if err2 == nil {
		buff := make([]byte, 10)
		for {
			cc2, remote2, rdErr2 := conn2.ReadFromUDP(buff)
			msg := string(buff[0:cc2])
			if rdErr2 != nil {
				fmt.Printf("net.ReadFromUDP() error2: %s\n", rdErr2)
			}
			if msg == "CONNECT" {
				_, _ = conn2.WriteTo([]byte("Connected"), remote2)
				fmt.Printf("A connection was received from client type B: \n")
				fmt.Printf("Remote address2: %v\n", remote2)
				go func(conn2 *net.UDPConn) {
					for {
						select {
						case <-quit:
							return
						default:
							if isAllowedClientAFlag == true {
								isAllowedClientAFlag = true
								cc, wrErr2 := conn2.WriteTo(packet[0:dataSizeForBClient], remote2)
								if wrErr2 != nil {
									fmt.Printf("net.WriteTo() error2: %s\n", wrErr2)
								} else {
									isAllowedClientAFlag = false
									fmt.Printf("Wrote %d bytes to client B type\n", cc)
									fmt.Printf("Bytes: %q\n", string(packet[:cc]))
									fmt.Printf("Remote address: %v\n", remote2)
								}
							}
						}
					}
					fmt.Printf("Out of infinite loop\n")
				}(conn2)
			} else {
				_, _ = conn2.WriteTo([]byte("NotConnected"), remote2)
				fmt.Println("Invalid command, Try again.")
			}
		}
	}
}
func BTypeConnection(outPortNumber int){
	conn2, err2 := net.ListenUDP("udp", &net.UDPAddr{IP:[]byte{0,0,0,0},Port:outPortNumber,Zone:""})
	if err2 != nil {
		log.Fatal(err2)
	}
	quit := make(chan struct{})
	listen(conn2, err2, quit)

}
func initPort(clientType string)(int){
	var portNumber int =0
	for {
		fmt.Printf("Please type in the port number that Client %v will connect with(within range 1024-65535): ",clientType)
		fmt.Scanf("%d\n", &portNumber)
		if isPortInRange(portNumber) {
			break
		}
	}
	return portNumber
}

func isPortInRange(portNumber int)(bool){
	if portNumber > 65535 || portNumber <1024 {
		fmt.Println("Invalid port, please try again.")
		return false
	}else{
		return true
	}
}

func readDataFromClientA(inPortNumber int){
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP:[]byte{0,0,0,0},Port:inPortNumber,Zone:""})
	if err != nil {
		log.Fatal(err)
	}

	for {
		dataSize, remote, rdErr := conn.ReadFromUDP(packet)
		dataSizeForBClient = dataSize
		fmt.Printf("New packet from client type A accepted: \n")
		isAllowedClientAFlag=true
		if rdErr != nil {
			fmt.Printf("net.ReadFromUDP() error: %s\n", rdErr)
		} else {
			fmt.Printf("Read %d bytes from socket\n", dataSize);
			fmt.Printf("Bytes: %q\n", string(packet[:dataSize]));
		}

		fmt.Printf("Remote address: %v\n", remote)
	}
	fmt.Printf("Out of infinite loop\n")
}

func run(){
	var inPortNumber = 0
	var outPortNumber = 0
	fmt.Println("Initializing Server...")
	inPortNumber = initPort("A")
	outPortNumber = initPort("B")
	fmt.Println("\nWaiting for incoming connections...\n")

	go BTypeConnection(outPortNumber)
	readDataFromClientA(inPortNumber)
}

func main() {
	run()
}
