package app

import "github.com/bitwormhole/starter/application"

func ExportConfigForExpressGoServerApp(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
