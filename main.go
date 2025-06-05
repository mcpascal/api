package main

import (
	"api/bootstraps"
	// _ "net/http/pprof"
)

func main() {
	app := bootstraps.NewApp()
	app.Start()
}
