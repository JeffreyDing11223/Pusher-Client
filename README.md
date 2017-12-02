# Pusher-Client

# example
# main.go

``` go 
package main

import (
	"github.com/JeffreyDing11223/Pusher-Client/pusher"
)

func main() {

	global.PusherCli= &pusher.PusherClient{
		Key: "de504dc5763aeef9ff52",
		PushUrl:"ws://ws.pusherapp.com:80/app/%s?protocol=7",
	}

	errpusher := global.PusherCli.New()
	if errpusher != nil {
		panic(errpusher)
	}
  
  ch:=make(chan int,1)
  <-ch

}
```
# pusher/connection.go
``` go
	var SubChannel []*Channel
	btcusd := p.Channel("")  //set channel what you want to receive
	ethusd := p.Channel("")  //set channel what you want to receive
	SubChannel=append(SubChannel,btcusd)
	SubChannel=append(SubChannel,ethusd)
	go Handler(SubChannel)
```
