package gate

import (
	"yfserver/game"
	"yfserver/msg"
)

func init() {
	msg.JsonProcessor.SetRouter(&msg.C2S_Login{}, game.ChanRpc)
}
