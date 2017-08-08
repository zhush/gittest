// yfclient project main.go
package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.19.11:13079")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		sendLogin(conn)
		wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		readMsg(conn)
		wg.Done()
	}()
	wg.Wait()
}

func sendLogin(conn net.Conn) {
	loginMsg := `{
		"C2S_Login":{
			"AccountName":"admin",
			"Password":"admin"
		}
	}`
	sendMsg(conn, loginMsg)
}

func sendMsg(conn net.Conn, msg string) {
	msgLen := len([]rune(msg))
	data := make([]byte, msgLen+2)
	binary.BigEndian.PutUint16(data, uint16(msgLen))
	copy(data[2:], []byte(msg))
	conn.Write(data)
}

func readMsgBuf(conn net.Conn) ([]byte, error) {
	headBuf := make([]byte, 2)
	if _, err := io.ReadFull(conn, headBuf); err != nil {
		return nil, err
	}
	msgLen := binary.BigEndian.Uint16(headBuf)
	recvBuf := make([]byte, msgLen)
	if _, err := io.ReadFull(conn, recvBuf); err != nil {
		return nil, err
	}
	return recvBuf, nil
}

func readMsg(conn net.Conn) {
	for {
		recvBuf, err := readMsgBuf(conn)
		if err != nil {
			fmt.Println("Receive error:", err.Error())
			break
		}
		handleRecvMsg(recvBuf)
	}
}

func handleRecvMsg(recvBuf []byte) {
	strMsg := string(recvBuf)
	fmt.Println("recvMsg:", strMsg)
}
