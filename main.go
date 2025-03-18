package main

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

func main() {

	var com_port string
	var port serial.Port

	for com_port != "//TerminalExit" {
		//Scan for ports and then display them for the user
		ports, err := serial.GetPortsList()
		if err != nil {
			log.Fatal(err)
		}
		if len(ports) == 0 {
			log.Fatal("No serial ports found!")
		}
		for _, port := range ports {
			fmt.Printf("Found port: %v\n", port)
		}

		fmt.Println("Type name of desired com port: ")
		fmt.Scan(&com_port)

		if com_port == "//TerminalExit" {
			return
		}

		//Open com port
		mode := &serial.Mode{
			BaudRate: 57600,
			Parity:   serial.EvenParity,
			DataBits: 7,
			StopBits: serial.OneStopBit,
		}

		port, err := serial.Open(com_port, mode)
		if err != nil {
			fmt.Println(port, ": Failed to open try again!")
		} else {
			break
		}
	}

	var Message string

	for Message != "//TerminalExit" {

		fmt.Scan(&Message)

		if Message == "//TerminalExit" {
			return
		}

		n, err := port.Write([]byte(Message))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Message contaning %v bytes sent to com port\n", n)

		buff := make([]byte, 100)
		for {
			n, err := port.Read(buff)
			if err != nil {
				log.Fatal(err)
				break
			}
			if n == 0 {
				fmt.Println("\nEOF")
				break
			}
			fmt.Printf("%v", string(buff[:n]))
		}
	}

}
