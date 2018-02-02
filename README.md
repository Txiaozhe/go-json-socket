# go-json-socket

#### A simple RPC tool, because RPC is needed to communicate with go services and node services in the encoding process, and [node-json-socket](https://github.com/sebastianseilund/node-json-socket) library is used for RPC communication between node services, so this library is written for better compatibility with node services.
#### 一个简易的rpc工具，由于编码过程中需要使用rpc进行go服务和node服务的通信，而node服务之间使用[node-json-socket](https://github.com/sebastianseilund/node-json-socket)库来进行rpc通信，为了更好地与node服务兼容，因此写了这个库。

## quick start

```shell
$ go get github.com/Txiaozhe/go-json-socket
```

```go
import (
	"log"
	jsonsocket "github.com/Txiaozhe/go-json-socket"
)

func main() {
	// connect to remote service.
	conn, err := jsonsocket.Connect("127.0.0.1:9838")
	if err != nil {
		log.Println("connect error: ", err)
	}

	token_info := Token{
		"Bearer eyJ1aWQiOjEyNTE4NDY3NDgsInR5cGUiOiJ0b29sIn0=.NWRhMWNlOTRiMTFjNmQwODM5YjA2Y2E5ZjZjMTBk",
	}
	auth_info := Auth{
		"tool",
		"auth_request",
		token_info,
	}

    // send a message to remote service.
	ch, err := jsonsocket.SendMessage(conn, auth_info)
	if err != nil {
		log.Println("send msg error: ", err)
	}

	log.Println(<-ch)

    // handle the message from remote service.
	ch1, err := jsonsocket.HandleMessage(conn)
	if err != nil {
		log.Println("handle msg error: ", err)
	}

	res := <-ch1
	log.Println(res, res.Len, res.Data)
}
```
