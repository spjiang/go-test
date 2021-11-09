
package main

import (
"log"
"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{
		Name: "COM45",
		Baud: 300,
		Size: 8,
	}
	//c := &serial.Config{
	//	Name: "COM45",
	//	Baud: 300,
	//	Size: 8,
	//	Parity: serial.ParityNone,
	//	StopBits: 1,
	//}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	//n, err := s.Write([]byte("test"))
	//if err != nil {
	//	log.Fatal(err)
	//}

	//保持数据持续接收
	for {
		buf := make([]byte, 1024)
		lens, err := s.Read(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		revData := buf[:lens]
		log.Printf("Rx:%X \n", revData)
	}
}
