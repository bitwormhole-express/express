package dev

import "github.com/bitwormhole/starter/application"

func ExportConfigForExpressGoServerDev(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
