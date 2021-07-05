package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		reader := bufio.NewReader(conn)
		for {
			slice, err := reader.ReadSlice('\n')
			if err != nil {
				continue
			}
			fmt.Printf("%s", slice)
		}

		////make a buffer
		//var buffer = bytes.Buffer{}
		//for {
		//	//read 10 byte every time
		//	data := make([]byte, 10)
		//	_, err := conn.Read(data)
		//	if err != nil {
		//		log.Printf("%s\n", err.Error())
		//		break
		//	}
		//	index := bytes.Index(data, []byte("\r\n"))
		//	//if found separator
		//	if index != -1 {
		//		buffer.Write(data[:index])
		//		receive := buffer.String()
		//		buffer.Reset() //reset buffer
		//		buffer.Write(data[index+2:])
		//
		//		log.Printf("rece msg: %s\n", receive)
		//		send := []byte(strings.ToUpper(receive) + "\r\n")
		//		_, err = conn.Write(send)
		//		if err != nil {
		//			log.Printf("send msg failed, error: %s\n", err.Error())
		//		}
		//		log.Printf("send msg: %s\n", receive)
		//	} else {
		//		//continue read from conn
		//		buffer.Write(data)
		//		continue
		//	}
		//}
	}
}
