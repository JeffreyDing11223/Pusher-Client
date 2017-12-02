package pusher

import (
	"fmt"
	"time"

	"golang.org/x/net/websocket"
)

type PusherClient struct{
	Key string
	PushUrl string
	conn     *websocket.Conn
	channels []*Channel
}


var PongMessgage = make(chan interface{} , 1)
var IfHasMessgage = make(chan interface{} , 1)
var goroutineStop bool= false
var countReconnect int
var countGoroutine int


func (p *PusherClient) New() (  er error) {

Reconnect_Loop:
	log.Println("connecting")
	ws, err := websocket.Dial(fmt.Sprintf(p.PushUrl, p.Key), "", "http://localhost/")
	if err != nil {
		logger.Warnning("dial err: ", err)
		time.Sleep(time.Second * 5)
		countReconnect+=1
		if countReconnect>1000{
			logger.Warnning("Reconnect too much") //重连次数过多
		}
		goto Reconnect_Loop
	}
	log.Println("connected"）
	countReconnect=0
	countGoroutine+=1
	//fmt.Println(runtime.NumGoroutine())
	if countGoroutine>1000{  //协程数过多
		log.Println("Goroutine too much")
	}
	p.conn=ws
	p.channels=[]*Channel{}
	goroutineStop=false
	go p.poll_pong()
	go p.ping()

	var SubChannel []*Channel
	btcusd := p.Channel("live_trades")
	ethusd := p.Channel("live_trades_ethusd")
	SubChannel=append(SubChannel,btcusd)
	SubChannel=append(SubChannel,ethusd)
	go Handler(SubChannel)
	return nil
}

func (p *PusherClient) ping() {

	ping := NewPingMessage()
	for {

		select {
		case <- IfHasMessgage:

		case <- time.After(120*time.Second): //120秒没有接受到消息则发送ping包
			{err:=websocket.JSON.Send(p.conn, ping)
			logger.Debug(ping)
			//fmt.Println(ping)
			if err!=nil{
				log.Println(err)
				return
			}
				select {
				case <- PongMessgage:

				case <- time.After(120*time.Second): //发送ping包后120秒没有收到pong包，则重连
					{go p.reconnect()
						return}
				}
			}
		}

	}
}

func (p *PusherClient) reconnect(){
	goroutineStop=true
	err:=p.Disconnect()
	if err!=nil{
		log.Println(err)
	}
	time.Sleep(60*time.Second)
	p.New()
}


func (p *PusherClient) poll_pong() {
	for {
		var msg Message
		err := websocket.JSON.Receive(p.conn, &msg)//阻塞
		if err != nil {
			log.Println(err)
			return
		}
		if msg.Event == "pusher:pong" {
			//fmt.Println(msg)
			logger.Debug(msg)
			PongMessgage<-msg
		}else if  msg.Event == "pusher:ping"{ //如果接受到server的ping包，则回应pong包
			
			err:=websocket.JSON.Send(p.conn, NewPongMessage())
			if err!=nil{
				log.Println(err)
				return
			}
		}else{
			IfHasMessgage<-msg
			p.processMessage(&msg)
		}

	}
}

func (p *PusherClient) processMessage(msg *Message) {
	for _, channel := range p.channels {
		if channel.Name == msg.Channel {
			channel.processMessage(msg)
		}
	}
}

func (p *PusherClient) Disconnect() error {
	return p.conn.Close()
}

func (p *PusherClient) Channel(name string) *Channel {
	for _, channel := range p.channels {
		if channel.Name == name {
			return channel
		}
	}
	channel := NewChannel(name)
	p.channels = append(p.channels, channel)
	websocket.JSON.Send(p.conn, NewSubscribeMessage(name))

	return channel
}
