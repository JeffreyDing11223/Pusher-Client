package pusher

type Channel struct {
	Name      string
	dataChans map[string][]chan interface{}
}

func NewChannel(name string) *Channel {
	return &Channel{name, make(map[string][]chan interface{})}
}

func (c *Channel) Bind(event string) chan interface{} {
	dataChan := make(chan interface{},10) //设置缓冲10个消息
	c.dataChans[event] = append(c.dataChans[event],dataChan)
	return dataChan
}

func (c *Channel) processMessage(msg *Message) {
	 value,ok:=c.dataChans[msg.Event]
	 if ok{
		 dataChan:=value[0]
		 dataChan <- msg.Data
	}

}
