package gojsonsocket

import (
	"bytes"
	"encoding/json"
	"log"
	"net"
	"regexp"
	"strconv"
	"time"
)

type Response struct {
	Len  int         `json:"len"`
	Data interface{} `json:"data"`
}

func Connect(addr string) (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", addr, 15*time.Second)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Listen(addr string) (net.Conn, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return nil, err
		}

		return conn, nil
	}
}

func SendMessage(conn net.Conn, d interface{}) (chan int, error) {
	ch := make(chan int)

	jsonData, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	delimeter := []byte("#")
	length := []byte(strconv.Itoa(len(jsonData)))

	data := bytesCombine(length, delimeter, jsonData)

	go func() {
		byteLen, err := conn.Write(data)
		if err != nil {
			log.Println("socket write error:", err)
			ch <- -1
		}

		ch <- byteLen
	}()

	return ch, nil
}

func HandleMessage(conn net.Conn) (chan Response, error) {
	ch := make(chan Response)

	buf := make([]byte, 1024)

	go func() {
		_, err := conn.Read(buf)
		if err != nil {
			ch <- Response{0, err}
		}

		reg := regexp.MustCompile(`(\S*)#(\S*)`)

		str := reg.FindStringSubmatch(string(buf))

		dataLen, err := strconv.Atoi(str[1])
		if err != nil {
			ch <- Response{0, err}
		}

		res := Response{dataLen, str[2]}
		ch <- res
	}()

	return ch, nil
}

func bytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
