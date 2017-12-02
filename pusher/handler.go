package pusher

func Handler(SubChannels []*Channel){
	subChan1:= SubChannels[0].Bind("trade")
	subchan2:= SubChannels[1].Bind("trade")

	for {
		if goroutineStop==true{
			return
		}

		select {
		case msg,ok:=<-subChan1:
			if ok{
				
				// do something
			}

		case msg,ok:=<-subchan2:
			if ok{
				
				// do something
			}

		case <-time.After(5*time.Second):
			continue
			
		}


	}
}
