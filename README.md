# go-json-socket

#### A simple RPC tool, because RPC is needed to communicate with go services and node services in the encoding process, and [node-json-socket](https://github.com/sebastianseilund/node-json-socket) library is used for RPC communication between node services, so this library is written for better compatibility with node services.
#### 一个简易的rpc工具，由于编码过程中需要使用rpc进行go服务和node服务的通信，而node服务之间使用[node-json-socket](https://github.com/sebastianseilund/node-json-socket)库来进行rpc通信，为了更好地与node服务兼容，因此写了这个库。

## quick start

```shell
$ go get github.com/Txiaozhe/go-json-socket
```

* go server
```go
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
```

* go client
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

#### this two go client and server can communicate with this two nodejs server and client.
#### 以上两个go实现的client和server可以与以下两个nodejs实现的client和server进行通信

* nodejs server

```shell
$ npm install json-socket --save
```

```javascript
const net = require('net');
const JsonSocket = require('json-socket');

const port = 3001;
const server = net.createServer();
server.listen(port);
server.on('connection', function(socket) {

  socket = new JsonSocket(socket);

  socket.on('message', function(message) {
    console.log('message');
    console.log(message);
  });

  socket.on('error', function(err) {
    console.log('error');
    console.log(err);
  });

  socket.sendMessage({code: 0, data: {mag: 'success'}}, err => {
    console.log(err);
  });
});
```

* nodejs client

```javascript
const net = require('net');
const JsonSocket = require('json-socket');

const port = 3001;
const host = '127.0.0.1';
const socket = new JsonSocket(new net.Socket());
socket.connect(port, host);
socket.on('connect', function() {
    socket.sendMessage({code: 0, data: {mag: 'success'}});
    socket.on('message', function(message) {
        console.log('The message is: ' + message);
    });
});
```
