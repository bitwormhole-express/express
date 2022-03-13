package main

import (
	communityserver "github.com/bitwomrhole-express/express/community-server"
	"github.com/bitwormhole/starter"
)

func main() {
	i := starter.InitApp()
	i.Use(communityserver.Module())
	i.Run()
}
