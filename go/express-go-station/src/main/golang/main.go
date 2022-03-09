package main

import (
	expressgoserver "bitwomrhole.com/djaf/express-go-server"
	"github.com/bitwormhole/starter"
)

func main() {
	i := starter.InitApp()
	i.Use(expressgoserver.Module())
	i.Run()
}
