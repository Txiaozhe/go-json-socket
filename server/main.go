package main

import (
	jsonsocket "github.com/Txiaozhe/go-json-socket"
	"log"
)

type Info struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

func main() {
	conn, err := jsonsocket.Listen("127.0.0.1:3001")
	if err != nil {
		log.Println("listen error: ", err)
	}

	ch, err := jsonsocket.HandleMessage(conn)
	if err != nil {
		log.Println("handel message error: ", err)
	}

	res := <-ch
	log.Println(res)

	msg := Info{
		0,
		"success",
	}
	ch1, err := jsonsocket.SendMessage(conn, msg)
	if err != nil {
		log.Println("response error:", err)
	}

	log.Println("send msg length: ", <-ch1)
}
