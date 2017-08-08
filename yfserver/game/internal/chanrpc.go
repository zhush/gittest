package internal

import (
	"fmt"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

var users = make(map[gate.Agent]struct{})

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	users[a] = struct{}{}
	fmt.Println("Client:", a.RemoteAddr(), " Connected!")

}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	delete(users, a)
	fmt.Println("Client:", a.RemoteAddr(), " disconnected!")
}

func broadcastMessage(msg []interface{}, _a gate.Agent) {
	for a := range users {
		if a != _a {
			a.WriteMsg(msg)
		}
	}
}
