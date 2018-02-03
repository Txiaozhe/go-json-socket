package main

import (
	jsonsocket "github.com/Txiaozhe/go-json-socket"
	"log"
)

type Token struct {
	Token string `json:"token"`
}

type Auth struct {
	Client     string `json:"client"`
	Event_type string `json:"type"`
	Data       Token  `json:"data"`
}

func main() {
	conn, err := jsonsocket.Connect("127.0.0.1:3001")
	if err != nil {
		log.Println("connect error: ", err)
	}

	token_info := Token{
		"Bearer eyJ1aWQiOjEyNTE4NDY3NDgsInR5cGUiOiJ0b29sIn0=.NWRhMWNlOTRiMTFjNmQwODM5YjA2Y2E5ZjZjMTBkZGQ0NDV",
	}
	auth_info := Auth{
		"tool",
		"auth_request",
		token_info,
	}

	ch, err := jsonsocket.SendMessage(conn, auth_info)
	if err != nil {
		log.Println("send msg error: ", err)
	}

	log.Println(<-ch)

	ch1, err := jsonsocket.HandleMessage(conn)
	if err != nil {
		log.Println("handle msg error: ", err)
	}

	res := <-ch1
	log.Println(res, res.Len, res.Data)
}
