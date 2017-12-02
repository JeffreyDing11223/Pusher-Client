# Pusher-Client

# example

``` go 
package main

import (
	"github.com/"
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
