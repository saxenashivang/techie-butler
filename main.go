package main

import (
	"github.com/saxenashivang/techiebutler/app"
	"github.com/saxenashivang/techiebutler/config"
)

func main() {
	config.LoadConfigs()
	// app entry point
	app.MapURL()
}
