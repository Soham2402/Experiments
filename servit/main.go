package main

import (
	"fmt"
	"log"
	"syscall"
)

const (
	IPv4 int = syscall.AF_INET
	STREAM int = syscall.SOCK_STREAM
	TCP int = syscall.IPPROTO_TCP
	DEFAULT_PORT int = 8080

)

var DEFAULT_IP = [4]byte{0,0,0,0}

func makeSocket()(syscall.Handle, error){
	fd, err := syscall.Socket(IPv4, STREAM, TCP)
	if err != nil { 
		return 0, err
	}
	return fd, nil	
}


// func bindSocket(fd *syscall.Handle, ip *[4]byte, port *int)error{
// 	sock_addr := syscall.SockaddrInet4{Port: port, Addr: ip}
// 	err  := syscall.Bind(fd, sock_addr)
// 	if err != nil{
// 		return err
// 	}
// 	return nil
// }

// func listenSocket()

func main() {
	soc, err := makeSocket()
	if err != nil {
		log.Fatal("There was an error while creating the port.", err)
	}
	// defer syscall.Close(soc)
	fmt.Println(soc)

	// err := bindSocket(&soc, &DEFAULT_IP, &DEFAULT_PORT)
	// if err != nil {
	// 	log.Fatal("There was an error while binding the port.", err)
	// }

	// err := listenSocket(&fd, &DEFAULT_BACKLOG)


}