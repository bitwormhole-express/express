package main

import (
	"github.com/bitwomrhole-express/express/station"
	"github.com/bitwormhole/starter"
)

func main() {
	i := starter.InitApp()
	i.Use(station.Module())
	i.Run()
}
