// yfserver project main.go
package main

import (
	"yfserver/conf"
	"yfserver/game"
	"yfserver/gate"

	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath

	leaf.Run(gate.Module,
		game.Module)
}
