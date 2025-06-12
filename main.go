package main

import (
	"api/bootstraps"
)

func main() {
	app := bootstraps.NewApp()
	app.Start()
}
