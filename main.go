//lint:file-ignore ST1006 heh...

package main

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
)


type SetState struct {
	value uint
}

type ResetState struct{}

type Handler struct {
	state uint
}

func NewHandler() actor.Receiver {
	return &Handler{};
}

func (self *Handler) Receive(ctx *actor.Context){
	switch msg := ctx.Message().(type){
	case actor.Initialized:{
		self.state = 10;
		fmt.Println("handler initialized, my state : ", self.state);
	}
	case actor.Started:{
		fmt.Println("handler started");
	}
	case SetState:{
		self.state = msg.value;
		fmt.Println("received new state", self.state);
	}
case ResetState:{
	self.state = 0;
	fmt.Println("resetting state, state's now : ", self.state)
}
	case actor.Stopped:{
		fmt.Println("handler stopped")
		_ = msg;
	}
	}
}

//go run --race main.go
func main(){
	engine := actor.NewEngine();
	processID := engine.Spawn(NewHandler, "handler");
	fmt.Println("processID -> ", processID);

	for i := 0; i < 10 ; i++ {
		go engine.Send(processID, SetState{ value : uint(i)})
	}
	engine.Send(processID, ResetState{})
}