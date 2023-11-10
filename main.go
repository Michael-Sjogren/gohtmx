package main

import (
	"htmxmain/htmx/model"
	"htmxmain/htmx/routes"
)

func main() {
	model.Setup()
	routes.SetupServerAndRun()
}
